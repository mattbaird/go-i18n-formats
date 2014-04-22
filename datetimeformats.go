package i18nformats

func GetDateTimeFormatFor(locale string) string {
	return GetDateFormatFor(locale) + " " + GetTimeFormatFor(locale)
}
