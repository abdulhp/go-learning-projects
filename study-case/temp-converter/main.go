package main

import (
	"fmt"
	"strings"
)

func main() {
	var menu string
	var exit bool
	
	fmt.Println("Temperature Converter")
	fmt.Println("Commands: convert, help, exit")
	for !exit {
		fmt.Print("> ")
		fmt.Scan(&menu)
		switch(menu) {
		case "convert":
			convert()
		case "help":
			help()
		case "exit":
			fmt.Println("Goodbye!")
			exit = true
		default:
			fmt.Printf("Input %v is not registered\n", menu)
		}	
	}
}

func convert() {
	var temp float32
	var srcUnit string
	var dstUnit string
	var tempResult float32

	fmt.Print("Enter temperature value(E.g. 32.1): ")
	fmt.Scan(&temp)
	fmt.Print("Enter source unit (C/F/K): ")
	fmt.Scan(&srcUnit)
	fmt.Print("Enter target unit (C/F/K): ")
	fmt.Scan(&dstUnit)

	srcUnit = strings.ToUpper(srcUnit)
	dstUnit = strings.ToUpper(dstUnit)	

	switch(srcUnit) {
		case "C":
			if dstUnit == "F" {
				tempResult = (temp * 9/5) + 32
				fmt.Printf("%.2f C = %.2f F\n", temp, tempResult)
			}else if dstUnit == "K" {
				tempResult = temp + 273.15
				fmt.Printf("%.2f C = %.2f K\n", temp, tempResult)
			}else {
				fmt.Println("Target unit cannot be the same as source unit")
			}
		case "F":
			if dstUnit == "C" {
				tempResult = (temp - 32) * 5/9
				fmt.Printf("%.2f F = %.2f C\n", temp, tempResult)
			}else if dstUnit == "K" {
				tempResult = (temp - 32) * 5/9 + 273.15
				fmt.Printf("%.2f F = %.2f K\n", temp, tempResult)
			} else {
				fmt.Println("Target unit cannot be the same as source unit")
			}
		case "K":
			if dstUnit == "C" {
				tempResult = temp - 273.15
				fmt.Printf("%.2f K = %.2f C\n", temp, tempResult)
			}else if dstUnit == "F" {
				tempResult = (temp - 273.15) * 9/5 + 32
				fmt.Printf("%.2f K = %.2f F\n", temp, tempResult)
			}else {
				fmt.Println("Target unit cannot be the same as source unit")
			}
		default:
			fmt.Printf("Unit %v is not registered", srcUnit)
		}
	}

func help() {
	fmt.Println("Usage:")
	fmt.Println("\tconvert\t- Start a new conversion")
	fmt.Println("\thelp\t- Show this help message")
	fmt.Println("\texit\t- Exit the program")
}
