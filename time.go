package timeformat

import (
	"strings"
	"time"
)

func Now(format string) string {
	currentTime := time.Now()
	parseFormat := ParseFormatGo(format)

	return currentTime.Format(parseFormat)
}

func Parse(datetime string, layout string, format string) (string, error) {
	layoutDateGo := ParseFormatGo(layout)
	formatDateGo := ParseFormatGo(format)
	result, err := time.Parse(layoutDateGo, datetime)
	if err != nil {
		return "", err
	}

	return result.Format(formatDateGo), nil
}

func ParseFormatGo(text string) string {
	text = strings.ReplaceAll(text, "hh", "15")
	text = strings.ReplaceAll(text, "h", "3")
	text = strings.ReplaceAll(text, "mm", "04")
	text = strings.ReplaceAll(text, "m", "4")
	text = strings.ReplaceAll(text, "sss", "05")
	text = strings.ReplaceAll(text, "ss", "05")
	text = strings.ReplaceAll(text, "s", "5")
	text = strings.ReplaceAll(text, "A", "PM")
	text = strings.ReplaceAll(text, "a", "pm")

	text = strings.ReplaceAll(text, "YYYY", "2006")
	text = strings.ReplaceAll(text, "YY", "06")
	text = strings.ReplaceAll(text, "MMMM", "January")
	text = strings.ReplaceAll(text, "MMM", "Jan")
	text = strings.ReplaceAll(text, "MM", "01")
	text = strings.ReplaceAll(text, "M", "1")
	text = strings.ReplaceAll(text, "DDDD", "Monday")
	text = strings.ReplaceAll(text, "DDD", "Mon")
	text = strings.ReplaceAll(text, "DD", "02")
	text = strings.ReplaceAll(text, "D", "2")

	return text
}
