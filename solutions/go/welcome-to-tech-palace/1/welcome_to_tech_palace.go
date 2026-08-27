package techpalace

import (
    "strings"
    "fmt"
)


// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	text := strings.ToUpper(customer)
	return "Welcome to the Tech Palace, " + text
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	star := strings.Repeat("*", numStarsPerLine)
	return fmt.Sprintf("%v\n%v\n%v", star, welcomeMsg, star)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	newText := strings.Replace(oldMsg, "*", " ", 999)
	res := strings.TrimSpace(newText)
	return res
}
