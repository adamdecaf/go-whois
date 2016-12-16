package whois

import (
	"fmt"
	"github.com/zonedb/zonedb"
	"io/ioutil"
	"net"
	"net/url"
	"strings"
	"time"
)

const (
	DEFAULT_WHOIS_PORT = "43"
	DEFAULT_TCP_TIMEOUT = time.Second * 10
)

// WhoisQuery accepts a domain and will return the full response of the whois query
// for the given domain. If the domain is invalid or whois fails an error will
// be returned.
func WhoisQuery(domain string) (string, error) {
	server, err := lookupWhoisServer(domain)
	if err != nil {
		return "", err
	}

	full_address := net.JoinHostPort(server, DEFAULT_WHOIS_PORT)
	conn, err := net.DialTimeout("tcp", full_address, DEFAULT_TCP_TIMEOUT)
	if err != nil {
		return "", err
	}

	payload := domain + "\r\n"
	conn.Write([]byte(payload))

	buffer, err := ioutil.ReadAll(conn)
	if err != nil {
		return "", err
	}

	conn.Close()

	resp := string(buffer[:])
	return resp, nil
}

// lookupWhoisServer digs into zonedb to find the whois server for the given
// tld. If no server is found (or domain is invalid) an error will be returned.
func lookupWhoisServer(domain string) (string, error) {
	_, err := url.Parse(domain)
	if err != nil {
		return "", err
	}

	split := strings.Split(domain, ".")
	// We assume there's at least a 'google' and 'com'
	if len(split) < 2 {
		return "", fmt.Errorf("Unable to parse domain (and find tld) for %s (split = %s)", domain, split)
	}

	zone := zonedb.PublicZone(domain)
	if zone == nil {
		return "", fmt.Errorf("unable to find zone for domain = %s", domain)
	}

	host := zone.WhoisServer()
	if host != "" {
		return host, nil
	}

	return "", fmt.Errorf("error finding whois server (empty host) for domain %s", domain)
}
