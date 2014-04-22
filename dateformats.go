package i18nformats

import (
	"strings"
)

// based on golang reference date/time layout: "Mon Jan 2 15:04:05 MST 2006"
// month 01
// day   02
// year  2006

const (
	// canada
	date_en_ca = "02/01/2006"
	date_fr_ca = "02/01/2006"
	// us
	date_en_us = "01/02/2006"
	// france
	date_fr = "02-01-2006"
	// china
	date_zh = "2006年01月02日"
)

func GetDateFormatFor(locale string) string {
	switch strings.ToLower(locale) {
	case "en_ca":
		return date_en_ca
	case "fr-ca", "fr_us":
		return date_fr_ca
	case "fr_BE", "fr_CA", "fr_CH", "fr_FR", "fr":
		return date_fr
	case "zh_CN", "zh_TW", "zh":
		return date_zh
	}
	return time_en_us
}
