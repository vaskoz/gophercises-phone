package phone

import "regexp"

var re = regexp.MustCompile(`\D`)

// Normalize removes all non-digit characters.
func Normalize(number string) string {
	return re.ReplaceAllString(number, "")
}
