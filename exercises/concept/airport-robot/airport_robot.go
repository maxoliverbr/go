package airportrobot

type Greeter interface {
	LanguageName() string
	Greet(g string) string
}

type German struct {
	Language string
}

type Italian struct {
	Language string
}

type Portuguese struct {
	Language string
}

func (German) LanguageName() string {
	return "German"
}

func (German) Greet(name string) string {
	return "Hallo " + name + "!"
}

func (Italian) LanguageName() string {
	return "Italian"
}

func (Italian) Greet(name string) string {
	return "I can speak Italian: Ciao " + name + "!"
}
func (Portuguese) LanguageName() string {
	return "Portuguese"
}

func (Portuguese) Greet(name string) string {
	return "I can speak Portuguese: Olá " + name + "!"
}

func SayHello(name string, g Greeter) string {
	return g.Greet(name)
}

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
