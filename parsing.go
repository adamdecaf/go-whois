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
	record := WhoisRecord{}
	rp := &record

	last_updated_at, err := find_last_updated_at(resp)
	if err == nil {
		// From standard: 'Mon Jan 2 15:04:05 MST 2006'
		// incoming: Fri Nov 20 07:42:26 GMT 2015
		t, err := time.Parse("Mon Jan 2 15:04:05 MST 2006", last_updated_at)
		if err == nil {
			rp.LastUpdatedAt = t
		}
	}

	created_at, err := find_created_at(resp)

	if err == nil {
		t, err := time.Parse("Mon Jan 2 15:04:05 MST 2006", created_at)
		if err == nil {
			rp.CreatedAt = t
		}
	}

	expires_at, err := find_expires_at(resp)

	if err == nil {
		t, err := time.Parse("Mon Jan 2 15:04:05 MST 2006", expires_at)
		if err == nil {
			rp.ExpiresAt = t
		}
	}

	return record, nil
}

func find_last_updated_at(resp string) (string, error) {
	r, err := regexp.Compile(`(?im)Updated Date: \s+(.+)$`)

	if err != nil {
		return "", err
	}

	return find_date_time(resp, r, "Updated Date")
}

func find_created_at(resp string) (string, error) {
	r, err := regexp.Compile(`(?im)Registration Date: \s+(.+)$`)

	if err != nil {
		return "", err
	}

	return find_date_time(resp, r, "Registration Date")
}

func find_expires_at(resp string) (string, error) {
	r, err := regexp.Compile(`(?im)Expiration Date: \s+(.+)$`)

	if err != nil {
		return "", err
	}

	return find_date_time(resp, r, "Expiration Date")
}

func find_date_time(resp string, r *regexp.Regexp, key string) (string, error) {
	loc := r.FindStringIndex(resp)

	if len(loc) != 2 {
		return "", fmt.Errorf("error getting location of match")
	}

	old := resp[loc[0]:loc[1]]
	// fmt.Printf("old = '%s'\n", old)

	match := strings.TrimSpace(strings.Replace(old, key + ":", "", -1))

	// fmt.Printf("match = '%s'\n", match)

	if match == "" {
		return "", fmt.Errorf("no matches found for '%s' record", key)
	}

	return match, nil
}
