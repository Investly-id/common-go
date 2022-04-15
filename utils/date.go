package utils

import (
	"fmt"
	"log"
	"time"
)

func StringToTime(str string, isDate bool) *time.Time {
	layout := "2006-01-02"
	if !isDate {
		layout = "2006-01-02 15:04:05"
	}
	t, err := time.Parse(layout, str)
	if err != nil {
		log.Println(err)
	}
	return &t
}

func FormatDateTime(t *time.Time, withTime bool) *string {

	if t == nil {
		return nil
	}

	var formattedText string

	if withTime {
		formattedText = t.Format("2006-01-02 15:04:05")
	} else {
		formattedText = t.Format("2006-01-02")
	}

	return &formattedText
}

func TimeSince(date *time.Time) string {

	currentTime := (time.Now().Unix() - date.Unix())

	interval := currentTime / 31536000
	if interval >= 1 {
		return fmt.Sprintf("%d tahun yang lalu", interval)
	}

	interval = currentTime / 2592000
	if interval >= 1 {
		return fmt.Sprintf("%d bulan yang lalu", interval)
	}

	interval = currentTime / 86400
	if interval >= 1 {
		return fmt.Sprintf("%d hari yang lalu", interval)
	}

	interval = currentTime / 3600
	if interval >= 1 {
		return fmt.Sprintf("%d jam yang lalu", interval)
	}

	interval = currentTime / 60
	if interval >= 1 {
		return fmt.Sprintf("%d menit yang lalu", interval)
	}
	return fmt.Sprintf("%d detik yang lalu", currentTime)
}
