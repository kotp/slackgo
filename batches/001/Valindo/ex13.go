package main

import "fmt"

type Speaker interface {
	sayHello() string
}

type English struct {
}

func (e English) sayHello() string {
	return "I speak English!"
}

type Chinese struct {
}

func (c Chinese) sayHello() string {
	return "I speak Mandarin!"
}

type Indian struct {
}

func (i Indian) sayHello() string {
	return "I speak Hindi"
}

func main() {
	speak := []Speaker{English{}, Chinese{}, Indian{}}

	for _, value := range speak {
		fmt.Println(value.sayHello())
	}
}
