package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(string) string
}

type Italian struct {
}

func (italian Italian) LanguageName() string {
	return "Italian"
}
func (italian Italian) Greet(name string) string {
	text := fmt.Sprintf("Ciao %v!", name)
	return text
}

type Portuguese struct {
}

func (portuguese Portuguese) LanguageName() string {
	return "Portuguese"
}

func (portuguese Portuguese) Greet(name string) string {
	text := fmt.Sprintf("Olá %v!", name)
	return text
}

func SayHello(name string, greeter Greeter) string {
	text := fmt.Sprintf("I can speak %v: %v", greeter.LanguageName(), greeter.Greet(name))
	return text
}