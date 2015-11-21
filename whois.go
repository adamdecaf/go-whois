package whois

import (
	"io/ioutil"
	"net"
	"time"
)

const (
	DEFAULT_WHOIS_PORT = "43"
	DEFAULT_TCP_TIMEOUT = time.Second * 10
)

func WhoisQuery(domain string) (string, error) {
	server, err := lookup_whois_server(domain)

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
