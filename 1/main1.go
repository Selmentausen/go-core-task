package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func main() {
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатеричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	DiscernTypes(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	concatenatedString := concatenateValues(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Printf("Concatenated values: %s\n", concatenatedString)
	runedString := []rune(concatenatedString)
	hashedString := hashWithSalt(runedString)
	fmt.Printf("Hashed values: %s\n", hashedString)
}

func DiscernTypes(vals ...interface{}) {
	for _, val := range vals {
		fmt.Printf("Type: %T\tValue: %v\n", val, val)
	}
}

func concatenateValues(vals ...interface{}) string {
	var sb strings.Builder
	for _, val := range vals {
		sb.WriteString(fmt.Sprintf("%v", val))
	}
	return sb.String()
}

func hashWithSalt(input []rune) string {
	mid := len(input) / 2
	salt := []rune("go-2024")
	var combined []rune
	combined = append(combined, input[:mid]...)
	combined = append(combined, salt...)
	combined = append(combined, input[mid:]...)
	hash := sha256.Sum256([]byte(string(combined)))
	return fmt.Sprintf("%x", hash)
}
