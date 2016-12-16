package whois

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

	last_updated_at, err := findLastUpdatedAt(resp)
	if err == nil {
		rp.LastUpdatedAt = last_updated_at
	}

	created_at, err := findCreatedAt(resp)
	if err == nil {
		rp.CreatedAt = created_at
	}

	expires_at, err := findExpiresAt(resp)
	if err == nil {
		rp.ExpiresAt = expires_at
	}

	registar_name, err := findRegisterName(resp)
	if err == nil {
		rp.Registrar = registar_name
	}

	registar_email, err := findRegisterEmail(resp)
	if err == nil {
		rp.ContactEmails = []string{registar_email}
	}

	return record, nil
}

func findRegisterName(blob string) (string, error) {
	patterns_and_formats := []*regexp.Regexp{
		regexp.MustCompile(`(?im)Registrant Name: (.+)$`),
		regexp.MustCompile(`(?im)Registrar Handle:(.+)$`),
		regexp.MustCompile(`(?im)Registrar:(.+)$`),
	}
	return findString(blob, patterns_and_formats, "ContactEmails")
}

func findRegisterEmail(blob string) (string, error) {
	patterns_and_formats := []*regexp.Regexp{
		regexp.MustCompile(`(?im)Registrant Email: (.+)$`),
	}
	return findString(blob, patterns_and_formats, "ContactEmails")
}

func findString(resp string, patterns []*regexp.Regexp, key string) (string, error) {
	for p := range patterns {
		res := patterns[p].FindStringSubmatch(resp)
		// Grab the first match
		if len(res) > 1 {
			return strings.TrimSpace(res[1]), nil
		}
	}
	return "", fmt.Errorf("unable to find patern for %s", key)
}

func findLastUpdatedAt(resp string) (time.Time, error) {
	patterns_and_formats := map[*regexp.Regexp]string{
		regexp.MustCompile(`(?im)Last Updated Date: \s+(.+)$`): "Mon Jan 2 15:04:05 MST 2006",
		regexp.MustCompile(`(?im)Last updated:(.+)$`): "2006-01-02",
		regexp.MustCompile(`(?im)Updated Date:\s+(.+)$`): "02-Jan-2006",
		regexp.MustCompile(`(?im)Updated Date:\s+(.+)$`): "2006-01-02T15:04:05Z",
	}
	return findDateTime(resp, patterns_and_formats, "LastUpdatedAt")
}

func findCreatedAt(resp string) (time.Time, error) {
	patterns_and_formats := map[*regexp.Regexp]string{
		regexp.MustCompile(`(?im)Registration Date: \s+(.+)$`): "Mon Jan 2 15:04:05 MST 2006",
		regexp.MustCompile(`(?im)Creation Date:\s+(.+)$`): "02-Jan-2006",
		regexp.MustCompile(`(?im)Creation Date:\s+(.+)$`): "2006-01-02T15:04:05Z",
		regexp.MustCompile(`(?im)Created: \s+(.+)$`): "2006-01-02",
	}
	return findDateTime(resp, patterns_and_formats, "CreatedAt")
}

func findExpiresAt(resp string) (time.Time, error) {
	patterns_and_formats := map[*regexp.Regexp]string{
		regexp.MustCompile(`(?im)Domain Expiration Date: \s+(.+)$`): "Mon Jan 2 15:04:05 MST 2006",
		regexp.MustCompile(`(?im)Expiration Date:\s+(.+)$`): "02-Jan-2006",
		regexp.MustCompile(`(?im)Registry Expiry Date:\s+(.+)$`): "2006-01-02T15:04:05Z",
	}
	return findDateTime(resp, patterns_and_formats, "ExpiresAt")
}

func findDateTime(resp string, patterns_and_formats map[*regexp.Regexp]string, key string) (time.Time, error) {
	for r, format := range patterns_and_formats {
		res := r.FindStringSubmatch(resp)

		// Grab the first match
		if len(res) > 1 {
			t, err := time.Parse(format, strings.TrimSpace(res[1]))
			if err == nil {
				return t, nil
			}
		}
	}
	return time.Now(), fmt.Errorf("unable to find patern for %s", key)
}
