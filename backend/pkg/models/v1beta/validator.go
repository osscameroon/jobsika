package v1beta

import (
	"regexp"
	"strings"
)

var (
	emailRegex = regexp.MustCompile(`(?i)[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}`)
	phoneRegex = regexp.MustCompile(`^[+]*[(]{0,1}[0-9]{1,4}[)]{0,1}[-\s\./0-9]*$`)
)

// IsEmailValid validates email string
func IsEmailValid(email string) bool {
	return emailRegex.MatchString(email)
}

// IsPhoneNumberValid validates phone number string
func IsPhoneNumberValid(phone string) bool {
	return phoneRegex.MatchString(phone)
}

// FormatTags formats tags string, to avoid duplicates, empty tags and leading/trailing spaces
func FormatTags(tagsString string) string {
	// tag should be separated by a comma,
	// each tag should be at least 1 characters long with no space,
	// tags with space will be ignored, tags should be unique
	if strings.TrimSpace(tagsString) != "" {
		tags := strings.Split(tagsString, ",")
		finalTags := ""
		finalTagsMap := make(map[string]struct{})
		for _, tag := range tags {
			tag = strings.TrimSpace(tag)
			if tag == "" {
				continue
			}

			if _, tagIsDuplicate := finalTagsMap[tag]; tagIsDuplicate {
				continue
			}

			if finalTags != "" {
				finalTags += ","
			}

			finalTags += tag

			finalTagsMap[tag] = struct{}{}
		}

		tagsString = finalTags
	}

	return tagsString
}
