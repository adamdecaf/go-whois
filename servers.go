package whois

import (
	"fmt"
	"github.com/zonedb/zonedb"
	"net/url"
	"strings"
)

func lookup_whois_server(domain string) (string, error) {
	_, err := url.Parse(domain)

	if err != nil {
		return "", err
	}

	split_host := strings.Split(domain, ".")

	// We assume there's at least a 'google' and 'com'
	if len(split_host) < 2 {
		return "", fmt.Errorf("Unable to parse domain (and find tld) for %s (split = %s)", domain, split_host)
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
