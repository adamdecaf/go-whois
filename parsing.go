package whois

import (
	"fmt"
	"time"
	"net/url"
	"regexp"
	"strings"
)

type WhoisRecord struct {
	Domain string

	// times
	LastUpdatedAt time.Time
	CreatedAt time.Time
	ExpiresAt time.Time

	// contact points
	Registrar string
	ContactEmails []string
	ContactPhoneNumbers []string

	// tech details
	NameServers []url.URL
	DNSSECEnabled bool
}

func ParseWhoisResponse(resp string) (WhoisRecord, error) {
	last_updated_at, err := find_last_updated_at(resp)
	record := WhoisRecord{}

	if err == nil {
		// From standard: 'Mon Jan 2 15:04:05 MST 2006'
		// incoming: Fri Nov 20 07:42:26 GMT 2015
		t, err := time.Parse("Mon Jan 2 15:04:05 MST 2006", last_updated_at)
		if err == nil {
			rp := &record
			rp.LastUpdatedAt = t
		}
	}

	return record, nil
}

func find_last_updated_at(resp string) (string, error) {
	r, err := regexp.Compile(`Updated Date: \s+([a-zA-Z0-9\s\:]{1,})\n`)

	if err != nil {
		return "", fmt.Errorf("unable to compile pattern due to %s", err)
	}

	loc := r.FindStringIndex(resp)

	if len(loc) != 2 {
		return "", fmt.Errorf("error getting location of match")
	}

	old := resp[loc[0]:loc[1]]
	// fmt.Printf("old = %s\n", old)

	match := strings.TrimSpace(strings.Replace(old, "Updated Date:", "", -1))

	// fmt.Printf("match = '%s'\n", match)

	if match == "" {
		return "", fmt.Errorf("no matches found for 'last updated at' record")
	}

	return match, nil
}
