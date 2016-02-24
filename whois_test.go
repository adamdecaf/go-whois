package whois

import (
	"fmt"
	"strings"
	"testing"
)

func TestSuccessfulWhoisLookup(t *testing.T) {
	resp, err := WhoisQuery("google.com")

	if err != nil {
		t.Errorf("error when finding whois server = %s", err)
	}

	if !strings.Contains(resp, "Whois Server Version 2.0") {
		t.Errorf("unable to validate whois response")
	}
}

func TestNonCOMWhois(t *testing.T) {
	resp, err := WhoisQuery("ashannon.us")

	if err != nil {
		t.Errorf("error getting whois lookup = %s", err)
	}

	if strings.Contains(resp, "Not found") {
		t.Errorf("error getting whois for ashannon.us")
	}

	if !strings.Contains(resp, "Adam Shannon") {
		t.Errorf("unable to properly get whois for ashannon.us, resp = %s", resp)
	}
}

func TestIgnoringSubdomains(t *testing.T) {
	resp, err := WhoisQuery("mail.google.com")

	if err != nil {
		t.Errorf("error trying to parse from subdomain = %s", err)
	}

	if !strings.Contains(resp, "Whois Server Version 2.0") {
		t.Errorf("unable to validate whois response")
	}
}

func TestFailingWhoisLookup(t *testing.T) {
	resp, err := WhoisQuery("no_tld")

	if err == nil {
		t.Errorf("we should have failed this whois lookup, but got %s", resp)
	}

	fmt.Printf("found expected error = %s\n", err)
}
