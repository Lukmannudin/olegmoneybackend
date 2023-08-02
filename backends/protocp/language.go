package protocp

import (
	"golang.org/x/text/language"
)

func LanguageTagToProto(t language.Tag) (string, error) {
	return t.String(), nil
}

func LanguageTagFromProto(s string) (language.Tag, error) {
	return language.Parse(s)
}
