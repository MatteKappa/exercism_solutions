package booking

import (
	"time"
	"strconv"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	data := date
	t, err := time.Parse(layout, data)
	if err != nil {
		return time.Now()
	}
	return t
	
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	today := time.Now()
	layout := "January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		return false
	}
	if t.Before(today) {
		return true
	}
	return false
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		return false
	}
	if t.Hour() >= 12  && t.Hour() < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		return "Error!"
	}
        formattedTime := t.Format("Monday, January 2, 2006, at 15:04")
	return "You have an appointment on " + formattedTime +"."
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	now := time.Now()
	currentY := strconv.Itoa(now.Year())
	anniversary := currentY + "-09-15"
	layout := "2006-01-02"

	t, err := time.Parse(layout, anniversary)
	if err != nil {
		return time.Now()
	}
	return t
}
