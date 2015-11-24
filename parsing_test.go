package whois

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestWhoisParserUS(t *testing.T) {
	resp, err := ioutil.ReadFile("test/ashannon.us")

	if err != nil {
		t.Errorf("error getting whois query response = %s", err)
	}

	parsed, err := ParseWhoisResponse(string(resp))

	if err != nil {
		t.Errorf("error parsing whois response = %s", err)
	}

	// print fields
	fmt.Printf("Domain = %s\n", parsed.Domain)

	fmt.Printf("LastUpdatedAt = %s\n", parsed.LastUpdatedAt)
	last_updated_str := parsed.LastUpdatedAt.String()
	if last_updated_str != "2015-11-20 07:42:26 +0000 GMT" {
		t.Errorf("last updated times don't match (actual = %s)", last_updated_str)
	}

	fmt.Printf("CreatedAt = %s\n", parsed.CreatedAt)
	created_at_str := parsed.CreatedAt.String()
	if created_at_str != "2008-12-21 20:11:01 +0000 GMT" {
		t.Errorf("created at times don't match (actual = %s)", created_at_str)
	}

	fmt.Printf("ExpiresAt = %s\n", parsed.ExpiresAt)
	expires_at_str := parsed.ExpiresAt.String()
	if expires_at_str != "2016-12-20 23:59:59 +0000 GMT" {
		t.Errorf("expires at times don't match (actual = %s)", expires_at_str)
	}

	fmt.Printf("Registrar = %s\n", parsed.Registrar)

	fmt.Printf("ContactEmails = %s\n", parsed.ContactEmails)

	fmt.Printf("ContactPhoneNumbers = %s\n", parsed.ContactPhoneNumbers)

	fmt.Printf("NameServers = %s\n", parsed.NameServers)

	fmt.Printf("DNSSECEnabled = %t\n", parsed.DNSSECEnabled)
}
