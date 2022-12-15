package utils

import "time"

func TimeInJKT(timeToConvert time.Time) time.Time {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}
	return timeToConvert.In(loc)
}

func NowInJKT() time.Time {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}
	return time.Now().In(loc)
}
