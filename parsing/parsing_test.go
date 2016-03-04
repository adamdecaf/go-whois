package parsing

import (
	"io/ioutil"
	"reflect"
	"testing"
)

// Uncomment and run to write a domain to a file
// func TestWriteFile(t *testing.T) {
// 	domain := "decafproductions.com"
// 	query, err := WhoisQuery(domain)

// 	if err != nil {
// 		panic(err)
// 	}

// 	ioutil.WriteFile("test/" + domain, []byte(query), 0755)
// }

type ParsedAnswer struct {
	// times
	LastUpdatedAt string
	CreatedAt string
	ExpiresAt string

	// contact points
	Registrar string
	ContactEmails []string
}

func TestWhoisParserUS(t *testing.T) {
	answer := ParsedAnswer{
		LastUpdatedAt: "2015-11-20 07:42:26 +0000 GMT",
		CreatedAt: "2008-12-21 20:11:01 +0000 GMT",
		ExpiresAt: "2016-12-20 23:59:59 +0000 GMT",
		Registrar: "Adam Shannon",
		ContactEmails: []string{"ashannon1000@gmail.com"},
	}
	VerifyParsedResposne("ashannon.us", answer, t)
}

func TestWhoisParserCOM(t *testing.T) {
	answer := ParsedAnswer{
		LastUpdatedAt: "2015-11-13 00:00:00 +0000 UTC",
		CreatedAt: "2012-12-13 00:00:00 +0000 UTC",
		ExpiresAt: "2016-12-13 00:00:00 +0000 UTC",
	}
	VerifyParsedResposne("decafproductions.com", answer, t)
}

func TestWhoisParserXOrg(t *testing.T) {
	answer := ParsedAnswer{
		LastUpdatedAt: "2007-01-12 21:42:24 +0000 UTC",
		CreatedAt: "1997-01-18 05:00:00 +0000 UTC",
		ExpiresAt: "2016-01-19 05:00:00 +0000 UTC",
		Registrar: "X.ORG Foundation, LLC",
		ContactEmails: []string{"leon@shiman.com"},
	}
	VerifyParsedResposne("x.org", answer, t)
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

	last_updated_str := parsed.LastUpdatedAt.String()
	if last_updated_str != answer.LastUpdatedAt {
		t.Errorf("last updated times don't match (actual = %s)", last_updated_str)
	}

	created_at_str := parsed.CreatedAt.String()
	if created_at_str != answer.CreatedAt {
		t.Errorf("created at times don't match (actual = %s)", created_at_str)
	}

	expires_at_str := parsed.ExpiresAt.String()
	if expires_at_str != answer.ExpiresAt {
		t.Errorf("expires at times don't match (actual = %s)", expires_at_str)
	}

	registrar_str := parsed.Registrar
	if registrar_str != answer.Registrar {
		t.Errorf("registrar doesn't match (actual = %s)", registrar_str)
	}

	contact_emails := parsed.ContactEmails
	if !reflect.DeepEqual(contact_emails, answer.ContactEmails) {
		t.Errorf("contact emails don't match (actual = %s)", contact_emails)
	}

	// fmt.Printf("Registrar = %s\n", parsed.Registrar)
	// fmt.Printf("ContactEmails = %s\n", parsed.ContactEmails)
	// fmt.Printf("ContactPhoneNumbers = %s\n", parsed.ContactPhoneNumbers)
	// fmt.Printf("NameServers = %s\n", parsed.NameServers)
	// fmt.Printf("DNSSECEnabled = %t\n", parsed.DNSSECEnabled)
}
