package booking

import (
	"time"
)

func Schedule(date string) (time.Time) {
	layout := "1/2/2006 15:04:05" 
	parsedTime, _ := time.Parse(layout, date)
	
	return parsedTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	nowTime := time.Now()
	parsedTime, _ := time.Parse("January 2, 2006 15:04:05", date) 
	if nowTime.After(parsedTime) {
		return true
	} else {
		return false
	}
	panic("")
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	startAft, _ := time.Parse(time.Kitchen,  "12:00PM")
	startAftHour := startAft.Hour()
	endAft, _ := time.Parse(time.Kitchen,  "6:00PM")
	endAftHour := endAft.Hour()
	parsedTime, _ := time.Parse("Monday, January 2, 2006 15:04:05", date) // Thursday, July 25, 2019 13:45:00
	parsedTimeHour := parsedTime.Hour()
	if parsedTimeHour >= startAftHour && parsedTimeHour < endAftHour {
		return true
	} else {
		return false
	}

}
// Description returns a formatted string of the appointment time.
func Description(date string) string {
	parsedTime, _ := time.Parse("1/2/2006 15:04:05", date) // 7/25/2019 13:45:00
	return parsedTime.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	t := time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
	return t
}
