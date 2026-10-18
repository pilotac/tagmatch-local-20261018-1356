// Package langtag normalises user-supplied language tags (for example "en_us" or "PT-br")
// into canonical BCP 47 form and picks the best match from a list of supported languages.
package langtag

import (
	"errors"
	"strings"

	"golang.org/x/text/language"
)

// ErrEmpty is returned when the input is blank.
var ErrEmpty = errors.New("langtag: empty tag")

// Normalize returns the canonical BCP 47 form of tag.
func Normalize(tag string) (string, error) {
	tag = strings.TrimSpace(strings.ReplaceAll(tag, "_", "-"))
	if tag == "" {
		return "", ErrEmpty
	}
	t, err := language.Parse(tag)
	if err != nil {
		return "", err
	}
	return t.String(), nil
}

// Best returns the supported language that best matches the Accept-Language header value.
func Best(acceptLanguage string, supported ...string) string {
	if len(supported) == 0 {
		return ""
	}
	tags := make([]language.Tag, 0, len(supported))
	for _, s := range supported {
		tags = append(tags, language.Make(s))
	}
	desired, _, err := language.ParseAcceptLanguage(acceptLanguage)
	if err != nil || len(desired) == 0 {
		return supported[0]
	}
	_, idx, _ := language.NewMatcher(tags).Match(desired...)
	return supported[idx]
}
