package main

import (
	"fmt"
	"os"
)

func convertToRoketinTime(hours, minutes, seconds int) (int, int, int) {
	totalSeconds := hours*3600 + minutes*60 + seconds

	earthSecondsPerDay := 24 * 60 * 60
	roketinSecondsPerDay := 10 * 100 * 100

	ratio := float64(roketinSecondsPerDay) / float64(earthSecondsPerDay)

	roketinTotalSeconds := int(float64(totalSeconds) * ratio)

	roketinHours := roketinTotalSeconds / (100 * 100)
	roketinMinutes := (roketinTotalSeconds % (100 * 100)) / 100
	roketinSeconds := roketinTotalSeconds % 100

	return roketinHours, roketinMinutes, roketinSeconds
}

func main() {
	var hours, minutes, seconds int

	fmt.Print("Enter hours (0-23): ")
	fmt.Scan(&hours)
	fmt.Print("Enter minutes (0-59): ")
	fmt.Scan(&minutes)
	fmt.Print("Enter seconds (0-59): ")
	fmt.Scan(&seconds)

	if hours < 0 || hours > 23 || minutes < 0 || minutes > 59 || seconds < 0 || seconds > 59 {
		fmt.Println("Invalid input! Hours must be 0-23, minutes and seconds must be 0-59")
		os.Exit(1)
	}

	roketinHours, roketinMinutes, roketinSeconds := convertToRoketinTime(hours, minutes, seconds)

	fmt.Printf("On Earth %02d:%02d:%02d, on planet Roketin Planet %02d:%02d:%02d\n", hours, minutes, seconds, roketinHours, roketinMinutes, roketinSeconds)
}
