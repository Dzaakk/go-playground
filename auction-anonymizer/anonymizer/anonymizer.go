package anonymizer

import (
	"fmt"
	"regexp"
)

type Anonymizer interface {
	Anonymize(input string) string
}

type EmailAnonymizer struct{}

func (e EmailAnonymizer) Anonymize(email string) string {
	regex := regexp.MustCompile(`(^[^@]{2})[^@]*(@.*$)`)
	// regex := regexp.MustCompile(`(^[^@]{2})[^@]*(@.*$)`)
	matches := regex.FindStringSubmatch(email)
	fmt.Printf("MATCHES = %v\n", matches)
	if len(matches) != 3 {
		return "****"
	}

	return matches[1] + "****" + matches[2]
}

type UsernameAnonymizer struct{}

func (u UsernameAnonymizer) Anonymize(username string) string {
	// regex := regexp.MustCompile(`^(.{3})(.*)$`)
	regex := regexp.MustCompile(`^(.{3})(.*)$`)
	return regex.ReplaceAllString(username, "${1}***")
}

type IPAnonymizer struct{}

func (i IPAnonymizer) Anonymize(ip string) string {
	// regex := regexp.MustCompile(`(\d+.\d+.\d+)\.\d+`)
	regex := regexp.MustCompile(`(\d+\.\d+\.\d+)\.\d+`)
	return regex.ReplaceAllString(ip, "${1}.***")
}
