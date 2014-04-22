package i18nformats

import (
	"strings"
)

// based on golang reference date/time layout: "Mon Jan 2 15:04:05 MST 2006"
// hour      15
// minute    04
// second    05
// time zone MST

const (
	// canada
	time_en_ca = "3:04PM"
	time_fr_ca = "3:04:05 PM"
	// us
	time_en_us = "3:04:05 PM"
	// france
	time_fr = "15:04:05"
	// china
	time_zh = "15:04:05"
)

func GetTimeFormatFor(locale string) string {
	switch strings.ToLower(locale) {
	case "en_ca":
		return time_en_ca
	case "fr-ca", "fr_us":
		return time_fr_ca
	case "fr_BE", "fr_CA", "fr_CH", "fr_FR", "fr":
		return time_fr
	case "zh_CN", "zh_TW", "zh":
		return time_zh
	}
	return time_en_us
}
