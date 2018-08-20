package i18n

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/ini.v1"
)

const fileExt = ".ini"

var i18n *languages
var ErrInvalidLocale = errors.New("invalid locale")

type languages struct {
	defaultLocale string
	translations  map[string]map[string]string
}

func (l *languages) init(translationsDir string, defaultLocale string) error {

	l.translations = make(map[string]map[string]string)
	if err := l.loadTranslations(translationsDir); err != nil {
		return err
	}

	if _, ok := l.translations[defaultLocale]; !ok {
		return errors.New("defaultLocale missing")
	}

	l.defaultLocale = defaultLocale

	return nil
}

func (l *languages) loadTranslations(translationsDir string) error {
	return filepath.Walk(translationsDir, func(filename string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fi.IsDir() {
			return nil
		}

		if strings.HasSuffix(strings.ToLower(fi.Name()), fileExt) {
			locale := strings.TrimSuffix(fi.Name(), fileExt)
			if err := l.verifyLoacle(locale); err != nil {
				return err
			}
			if message, err := ini.LoadSources(ini.LoadOptions{
				IgnoreInlineComment:         true,
				UnescapeValueCommentSymbols: true,
			}, filename); err != nil {
				return err
			} else {
				for _, section := range message.Sections() {
					if sectionName := section.Name(); sectionName == "DEFAULT" {
						l.translations[locale] = section.KeysHash()
					} else {
						for k, v := range section.KeysHash() {
							l.translations[locale][sectionName+"."+k] = v
						}
					}
				}
			}
		}

		return nil
	})
}

func (l languages) verifyLoacle(locale string) error {
	parts := strings.Split(locale, "_")
	if len(parts) != 2 {
		return ErrInvalidLocale
	}

	if len(parts[0]) != 2 {
		return ErrInvalidLocale
	}

	for _, r := range parts[0] {
		if r < rune('a') || r > rune('z') {
			return ErrInvalidLocale
		}
	}

	if len(parts[1]) != 2 {
		return ErrInvalidLocale
	}

	for _, r := range parts[1] {
		if r < rune('A') || r > rune('Z') {
			return ErrInvalidLocale
		}
	}

	return nil
}

func Init(translationsDir string, defaultLocale string) error {
	i18n = &languages{}
	return i18n.init(translationsDir, defaultLocale)
}

func Tr(locale string, key string, args ...interface{}) string {
	if i18n != nil {
		if locale == "" {
			locale = i18n.defaultLocale
		}

		if translations, ok := i18n.translations[locale]; ok {
			if format, ok := translations[key]; ok {
				if args != nil {
					return fmt.Sprintf(format, args...)
				} else {
					return format
				}
			}
		}

		if locale != i18n.defaultLocale {
			if format, ok := i18n.translations[i18n.defaultLocale][key]; ok {
				if args != nil {
					return fmt.Sprintf(format, args...)
				} else {
					return format
				}
			}
		}
	}
	return key
}
