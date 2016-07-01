package parsing

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type WhoisRecord struct {
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
		rp.LastUpdatedAt = last_updated_at
	}

	created_at, err := find_created_at(resp)
	if err == nil {
		rp.CreatedAt = created_at
	}

	expires_at, err := find_expires_at(resp)
	if err == nil {
		rp.ExpiresAt = expires_at
	}

	registar_name, err := find_registar_name(resp)
	if err == nil {
		rp.Registrar = registar_name
	}

	registar_email, err := find_registar_email(resp)
	if err == nil {
		rp.ContactEmails = []string{registar_email}
	}

	return record, nil
}

func find_registar_name(blob string) (string, error) {
	patterns_and_formats := []*regexp.Regexp{
		regexp.MustCompile(`(?im)Registrant Name: (.+)$`),
	}
	return find_string(blob, patterns_and_formats, "ContactEmails")
}

func find_registar_email(blob string) (string, error) {
	patterns_and_formats := []*regexp.Regexp{
		regexp.MustCompile(`(?im)Registrant Email: (.+)$`),
	}
	return find_string(blob, patterns_and_formats, "ContactEmails")
}

func find_string(resp string, patterns []*regexp.Regexp, key string) (string, error) {
	for p := range patterns {
		res := patterns[p].FindStringSubmatch(resp)
		// Grab the first match
		if len(res) > 1 {
			return strings.TrimSpace(res[1]), nil
		}
	}
	return "", fmt.Errorf("unable to find patern for %s", key)
}

func find_last_updated_at(resp string) (time.Time, error) {
	patterns_and_formats := map[*regexp.Regexp]string{
		regexp.MustCompile(`(?im)Last Updated Date: \s+(.+)$`): "Mon Jan 2 15:04:05 MST 2006",
		regexp.MustCompile(`(?im)Updated Date:\s+(.+)$`): "02-Jan-2006",
		regexp.MustCompile(`(?im)Updated Date:\s+(.+)$`): "2006-01-02T15:04:05Z",
	}
	return find_date_time(resp, patterns_and_formats, "LastUpdatedAt")
}

func find_created_at(resp string) (time.Time, error) {
	patterns_and_formats := map[*regexp.Regexp]string{
		regexp.MustCompile(`(?im)Registration Date: \s+(.+)$`): "Mon Jan 2 15:04:05 MST 2006",
		regexp.MustCompile(`(?im)Creation Date:\s+(.+)$`): "02-Jan-2006",
		regexp.MustCompile(`(?im)Creation Date:\s+(.+)$`): "2006-01-02T15:04:05Z",
	}
	return find_date_time(resp, patterns_and_formats, "CreatedAt")
}

func find_expires_at(resp string) (time.Time, error) {
	patterns_and_formats := map[*regexp.Regexp]string{
		regexp.MustCompile(`(?im)Domain Expiration Date: \s+(.+)$`): "Mon Jan 2 15:04:05 MST 2006",
		regexp.MustCompile(`(?im)Expiration Date:\s+(.+)$`): "02-Jan-2006",
		regexp.MustCompile(`(?im)Registry Expiry Date:\s+(.+)$`): "2006-01-02T15:04:05Z",
	}
	return find_date_time(resp, patterns_and_formats, "ExpiresAt")
}

func find_date_time(resp string, patterns_and_formats map[*regexp.Regexp]string, key string) (time.Time, error) {
	for r, format := range patterns_and_formats {
		res := r.FindStringSubmatch(resp)

		// Grab the first match
		if len(res) > 1 {
			t, err := time.Parse(format, res[1])
			if err == nil {
				return t, nil
			}
		}
	}
	return time.Now(), fmt.Errorf("unable to find patern for %s", key)
}
