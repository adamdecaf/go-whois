package whois

import (
	"fmt"
	"io/ioutil"
	"testing"
)

type ParsedAnswer struct {
	// times
	LastUpdatedAt string
	CreatedAt string
	ExpiresAt string
}

func TestWhoisParserUS(t *testing.T) {
	answer := ParsedAnswer{
		LastUpdatedAt: "2015-11-20 07:42:26 +0000 GMT",
		CreatedAt: "2008-12-21 20:11:01 +0000 GMT",
		ExpiresAt: "2016-12-20 23:59:59 +0000 GMT",
	}
	VerifyParsedResposne("ashannon.us", answer, t)
}

func VerifyParsedResposne(domain string, answer ParsedAnswer, t *testing.T) {
	resp, err := ioutil.ReadFile("test/" + domain)

	if err != nil {
		t.Errorf("error getting whois query response = %s", err)
	}

	parsed, err := ParseWhoisResponse(string(resp))

	if err != nil {
		t.Errorf("error parsing whois response = %s", err)
	}

	fmt.Printf("LastUpdatedAt = %s\n", parsed.LastUpdatedAt)
	last_updated_str := parsed.LastUpdatedAt.String()
	if last_updated_str != answer.LastUpdatedAt {
		t.Errorf("last updated times don't match (actual = %s)", last_updated_str)
	}

	fmt.Printf("CreatedAt = %s\n", parsed.CreatedAt)
	created_at_str := parsed.CreatedAt.String()
	if created_at_str != answer.CreatedAt {
		t.Errorf("created at times don't match (actual = %s)", created_at_str)
	}

	fmt.Printf("ExpiresAt = %s\n", parsed.ExpiresAt)
	expires_at_str := parsed.ExpiresAt.String()
	if expires_at_str != answer.ExpiresAt {
		t.Errorf("expires at times don't match (actual = %s)", expires_at_str)
	}

	// fmt.Printf("Registrar = %s\n", parsed.Registrar)
	// fmt.Printf("ContactEmails = %s\n", parsed.ContactEmails)
	// fmt.Printf("ContactPhoneNumbers = %s\n", parsed.ContactPhoneNumbers)
	// fmt.Printf("NameServers = %s\n", parsed.NameServers)
	// fmt.Printf("DNSSECEnabled = %t\n", parsed.DNSSECEnabled)
}
