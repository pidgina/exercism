package booking

import (
    "time"
"fmt"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"

	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Ошибка парсинга:", err)
		return time.Time{}
	}
	return t

}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	now := time.Now()
	layout := "January 2, 2006 15:04:05"

	tDate, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Ошибка парсинга в функции вычитания:", err)
		return false
	}

return now.After(tDate)

}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	tDate, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Ошибка парсинга в функции /половина дня/:", err)
		return false
	}

	if tDate.Hour() >= 12 && tDate.Hour() < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:4:05"

	tDate, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Ошибка парсинга в функции Description:", err)
		return ""
	}
	formLayout := "You have an appointment on Monday, January 2, 2006, at 15:4."
	res := tDate.Format(formLayout)
	return res
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	open := time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)

	return open
}
