package utils

import (
	"log"
	"math"
	"strconv"
	"time"
)

func ParseDateTimeInJakarta(timeStr string) (*time.Time, error) {
	timeLocation, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return nil, err
	}

	parsedTime, err := time.ParseInLocation("2006-01-02T15:04:05.999Z", timeStr, timeLocation)
	if err != nil {
		return nil, err
	}

	return &parsedTime, nil
}

func ParseBool(str string) bool {
	boolValue, err := strconv.ParseBool(str)
	if err != nil {
		log.Fatal(err)
	}

	return boolValue
}

func ParseNullString(value string) *string {
	if len(value) == 0 {
		return nil
	}

	return &value
}

func ParseInt64(value string) *int64 {
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil
	}

	return &i
}

func ParseInt16(value string) *int16 {
	i, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		return nil
	}
	i16 := int16(i)

	return &i16
}

func ParseInt(value string) *int {
	i, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return nil
	}
	i0 := int(i)

	return &(i0)
}

func ParseFloat32(value string) *float32 {
	f, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return nil
	}
	f32 := float32(f)

	return &f32
}

func ParseExcelTimeToTime(excelTime string) *time.Time {
	timeF64, err := strconv.ParseFloat(excelTime, 64)
	if err != nil {
		return nil
	}

	time := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC).Add(time.Duration(math.Round(timeF64 * float64(24*time.Hour))))

	return &time
}

func ParseExcelDateToTime(excelDate string) *time.Time {
	date, err := time.Parse("01-02-06", excelDate)
	if err != nil {
		return nil
	}

	return &date
}
