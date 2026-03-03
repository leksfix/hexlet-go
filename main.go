package main

import (
	"fmt"
	"hexlet-go/greeting"
	"github.com/fatih/color"
)

func main () {
	fmt.Println(greeting.Hello())
	color.Red(greeting.Hello())
	c := color.New(color.BgGreen, color.FgYellow)
	c.Println("TEST TEST TEST!")
}
