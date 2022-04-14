package utils

import "time"

func FormatDateTime(t *time.Time, withTime bool) string {
	if withTime {
		return t.Format("2006-01-02 15:04:05")
	}
	return t.Format("2006-01-02")
}
