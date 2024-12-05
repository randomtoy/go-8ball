package translator

import (
	"github.com/eduardolat/goeasyi18n"
)

func initialize(code, value string) (*goeasyi18n.I18n, error) {
	i18n := goeasyi18n.NewI18n()

	translation, err := goeasyi18n.LoadFromJsonFiles("./_resources/translation/" + code + "/*.json")
	if err != nil {
		return nil, err
	}

	i18n.AddLanguage(code, translation)

	return i18n, nil
}

func GetStringByCode(code, value string) (string, error) {
	i18n, err := initialize(code, value)
	if err != nil {
		return "", err
	}

	opts := goeasyi18n.Options{}

	return i18n.T(code, value, opts), nil
}
