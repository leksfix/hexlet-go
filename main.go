package main

import (
	"fmt"
	"github.com/fatih/color"
	"hexlet-go/greeting"
	"hexlet-go/employees"
	"hexlet-go/payroll"
)

func main() {
	fmt.Println(greeting.Hello())
	color.Red(greeting.Hello())
	c := color.New(color.BgGreen, color.FgYellow)
	c.Println("Hexlet for Brave!")

	e1 := employees.Employee{Name: "Emp1", Position: "pos1", Skills: []string{"tup", "glup"}}
	e2 := employees.Employee{Name: "Emp2", Position: "pos2", Skills: []string{"aaa", "bbb"}}
	
	//e1.baseSalary = 600

	e1.SetBaseSalary(600)
	e2.SetBaseSalary(700)

	gross1 := payroll.CalcGross(e1, 11)
	net2 := payroll.CalcNet(800, 13)

	fmt.Printf("Gross 1: %.2f\n", gross1)
	fmt.Printf("Net 2: %.3f\n", net2)
}
