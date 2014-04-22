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
	case "fr-ca", "fr_ca", "fr_us", "fr-us":
		return time_fr_ca
	case "fr-be", "fr-ch", "fr-fr", "fr_be", "fr_ch", "fr_fr", "fr":
		return time_fr
	case "zh-cn", "zh-tw", "zh_cn", "zh_tw", "zh":
		return time_zh
	}
	return time_en_us
}
