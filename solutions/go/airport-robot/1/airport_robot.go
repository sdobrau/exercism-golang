package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
	LanguageName() string
	Greet(visitor string) string
	
}

func SayHello(name string, g Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(name))
}

type Italian struct {
	Language string
}

func (i Italian) Greet(visitor string) string {
	return "Ciao " + visitor + "!"
}

func (i Italian) LanguageName() string {
	return "Italian"
}

type Portuguese struct {
	Language string
}

func (p Portuguese) Greet(visitor string) string {
	return "Olá " + visitor + "!"
}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}
