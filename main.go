package main

import (
	"fmt"
	"github.com/fatih/color"
	"hexlet-go/greeting"
)

func main() {
	fmt.Println(greeting.Hello())
	color.Red(greeting.Hello())
	c := color.New(color.BgGreen, color.FgYellow)
	c.Println("TEST TEST TEST!")
}
