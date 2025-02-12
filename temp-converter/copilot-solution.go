package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	kelvinOffset    = 273.15
	fahrenheitBase  = 32.0
	fahrenheitScale = 9.0 / 5.0
)

type tempUnit string

const (
	Celsius    tempUnit = "C"
	Fahrenheit tempUnit = "F"
	Kelvin     tempUnit = "K"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("Temperature Converter")
	fmt.Println("Commands: convert, help, exit")

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		command := strings.TrimSpace(input)

		switch command {
		case "convert":
			convert(reader)
		case "help":
			help()
		case "exit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Printf("Command '%s' is not recognized\n", command)
		}
	}
}

func convert(reader *bufio.Reader) {
	// Get temperature value
	var temp float64
	for {
		fmt.Print("Enter temperature value: ")
		_, err := fmt.Fscanf(reader, "%f\n", &temp)
		if err == nil {
			break
		}
		fmt.Println("Please enter a valid number")
		// Clear the invalid input
		reader.ReadString('\n')
	}

	// Get source unit
	srcUnit := getUnit(reader, "Enter source unit (C/F/K): ")
	if srcUnit == "" {
		return
	}

	// Get target unit
	dstUnit := getUnit(reader, "Enter target unit (C/F/K): ")
	if dstUnit == "" {
		return
	}

	if srcUnit == dstUnit {
		fmt.Println("Source and target units cannot be the same")
		return
	}

	result := convertTemperature(temp, tempUnit(srcUnit), tempUnit(dstUnit))
	fmt.Printf("%.2f°%s = %.2f°%s\n", temp, srcUnit, result, dstUnit)
}

func getUnit(reader *bufio.Reader, prompt string) string {
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return ""
		}

		unit := strings.ToUpper(strings.TrimSpace(input))
		if unit == "C" || unit == "F" || unit == "K" {
			return unit
		}
		fmt.Println("Please enter C, F, or K")
	}
}

func convertTemperature(temp float64, from, to tempUnit) float64 {
	// First convert to Celsius as intermediate step
	var celsius float64
	switch from {
	case Celsius:
		celsius = temp
	case Fahrenheit:
		celsius = (temp - fahrenheitBase) / fahrenheitScale
	case Kelvin:
		celsius = temp - kelvinOffset
	}

	// Then convert from Celsius to target unit
	switch to {
	case Celsius:
		return celsius
	case Fahrenheit:
		return celsius*fahrenheitScale + fahrenheitBase
	case Kelvin:
		return celsius + kelvinOffset
	default:
		return 0
	}
}

func help() {
	fmt.Println("Usage:")
	fmt.Println("  convert - Start a new conversion")
	fmt.Println("  help    - Show this help message")
	fmt.Println("  exit    - Exit the program")
}
