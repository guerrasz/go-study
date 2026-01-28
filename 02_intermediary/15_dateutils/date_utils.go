package main

import (
	"fmt"
	"time"
)

func main() {
	// get current time
	t := time.Now()

	fmt.Printf("%v\n", t)

	// custom format
	fmt.Printf("%v\n", t.Format("02/01/2006 15:04:05"))

	// add time
	t = t.Add(time.Hour * 24)
	fmt.Printf("%v\n", t)

	// sub time
	t = t.Add(-time.Hour * 24)
	fmt.Printf("%v\n", t)

	// specific time
	specificTime := time.Date(2025, time.January, 30, 12, 0, 0, 0, time.UTC)
	fmt.Printf("%v\n", specificTime)

	// parsed time (layout, actual date)
	// base layou date is 02/01/2006 15:04:05
	parsedTime, err := time.Parse("2006-01-02", "2025-01-30")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", parsedTime)

	// time in different TZ
	loc, _ := time.LoadLocation("Asia/Tokyo")
	tTokyo := t.In(loc)
	fmt.Printf("%v\n", tTokyo)

	fmt.Printf("Truncated time in tokyo: %v\n", tTokyo.Truncate(time.Hour))
	fmt.Printf("Rounded time in tokyo: %v\n", tTokyo.Round(time.Hour))

	// sub 2 times to get diff
	fmt.Printf("Diff: %v\n", t.Sub(tTokyo))

	// compare 2 times with after and before
	fmt.Printf("After: %v\n", t.After(tTokyo))
	fmt.Printf("Before: %v\n", t.Before(tTokyo))
	fmt.Printf("Equal: %v\n", t.Equal(tTokyo))

	// print unix time
	fmt.Printf("Unix time: %v\n", t.Unix())
	fmt.Printf("Unix time in tokyo: %v\n", tTokyo.Unix())
}
