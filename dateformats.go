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
	case "fr-ca", "fr_ca", "fr_us", "fr-us":
		return date_fr_ca
	case "fr-be", "fr-ch", "fr-fr", "fr_be", "fr_ch", "fr_fr", "fr":
		return date_fr
	case "zh-cn", "zh-tw", "zh_cn", "zh_tw", "zh":
		return date_zh
	}
	return date_en_us
}
