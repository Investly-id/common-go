package utils_test

import (
	"testing"
	"time"

	"github.com/Investly-id/common-go/v2/utils"
	"github.com/stretchr/testify/assert"
)

func Test_StringToTime_Date(t *testing.T) {
	time := utils.StringToTime("2020-10-02", true)
	assert.Equal(t, 2020, time.Year())
	assert.Equal(t, 10, int(time.Month()))
	assert.Equal(t, 2, time.Day())
}

func Test_StringToTime_DateTime(t *testing.T) {
	time := utils.StringToTime("2020-10-02 12:12:12", false)
	assert.Equal(t, 2020, time.Year())
	assert.Equal(t, 10, int(time.Month()))
	assert.Equal(t, 2, time.Day())
	assert.Equal(t, 12, time.Hour())
	assert.Equal(t, 12, time.Minute())
	assert.Equal(t, 12, time.Second())
}

func Test_FormaDateTime_WithTimeTrue(t *testing.T) {
	time, _ := time.Parse("2006-01-02 15:04:05", "2020-10-10 10:10:10")
	formattedDate := utils.FormatDateTime(&time, false)
	assert.Equal(t, "2020-10-10", *formattedDate)
}

func Test_FormaDateTime_WithTimeFalse(t *testing.T) {
	time, _ := time.Parse("2006-01-02 15:04:05", "2020-10-10 10:10:10")
	formattedDate := utils.FormatDateTime(&time, true)
	assert.Equal(t, "2020-10-10 10:10:10", *formattedDate)
}
