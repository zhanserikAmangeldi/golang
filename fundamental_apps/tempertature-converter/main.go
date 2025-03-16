package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Options:\n1. From Celsius to Fahrenheit\n2. From Fahrenheit to Celsius")
	var option string
	var temperature string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter option: ")
	option, _ = reader.ReadString('\n')

	fmt.Println("Enter temperature: ")
	temperature, _ = reader.ReadString('\n')

	temp, err := strconv.ParseFloat(strings.TrimSpace(temperature), 64)
	if err != nil {
		fmt.Println("Invalid temperature")
		return
	}

	if strings.TrimSpace(option) == "1" {
		fmt.Println(celsiusToFahrenheit(temp))
	} else if strings.TrimSpace(option) == "2" {
		fmt.Println(fahrenheitToCelsius(temp))
	} else {
		fmt.Println("Invalid option")
	}
}

func celsiusToFahrenheit(temperature float64) float64 {
	return temperature*9/5 + 32
}

func fahrenheitToCelsius(temperature float64) float64 {
	return (temperature - 32) * 5 / 9
}
