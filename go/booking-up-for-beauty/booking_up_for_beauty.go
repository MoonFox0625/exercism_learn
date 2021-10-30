package booking

import (
	"log"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

// HasPassed returns whether a date has passed
// "July 25, 2019 13:45:00"
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		log.Fatal(err)
	}
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
// "Thursday, July 25, 2019 13:45:00"
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		log.Fatal(err)
	}
	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time
// Input: "7/25/2019 13:45:00"
// Output: "You have an appointment on Thursday, July 25, 2019, at 13:45."
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		log.Fatal(err)
	}
	return t.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
