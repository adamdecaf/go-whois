package whois

import (
	"fmt"
	"testing"
)

func TestWhoisParserUS(t *testing.T) {
	resp, err := WhoisQuery("ashannon.us")

	if err != nil {
		t.Errorf("error getting whois query response = %s", err)
	}

	parsed, err := ParseWhoisResponse(resp)

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

	fmt.Printf("ExpiresAt = %s\n", parsed.ExpiresAt)

	fmt.Printf("Registrar = %s\n", parsed.Registrar)

	fmt.Printf("ContactEmails = %s\n", parsed.ContactEmails)

	fmt.Printf("ContactPhoneNumbers = %s\n", parsed.ContactPhoneNumbers)

	fmt.Printf("NameServers = %s\n", parsed.NameServers)

	fmt.Printf("DNSSECEnabled = %t\n", parsed.DNSSECEnabled)
}
