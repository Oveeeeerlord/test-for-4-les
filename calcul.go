package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calculator struct {
	operand1 int
	operand2 int
	operator string
}

var romanToArabic = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func (c *Calculator) Add() int {
	return c.operand1 + c.operand2
}

func (c *Calculator) Subtract() int {
	return c.operand1 - c.operand2
}

func (c *Calculator) Multiply() int {
	return c.operand1 * c.operand2
}

func (c *Calculator) Divide() int {
	if c.operand2 == 0 {
		panic("Деление на ноль")
	}
	return c.operand1 / c.operand2
}

func IsRomanNumeral(s string) bool {
	for _, c := range s {
		if _, ok := romanToArabic[string(c)]; !ok {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите первое число: ")
	operand1Str, _ := reader.ReadString('\n')
	operand1, _ := strconv.Atoi(strings.TrimSpace(operand1Str))

	fmt.Print("Введите оператор (+, -, *, /): ")
	operator, _ := reader.ReadString('\n')
	operator = strings.TrimSpace(operator)

	fmt.Print("Введите второе число: ")
	operand2Str, _ := reader.ReadString('\n')
	operand2, _ := strconv.Atoi(strings.TrimSpace(operand2Str))

	calculator := &Calculator{
		operand1: operand1,
		operator: operator,
		operand2: operand2,
	}

	var result int

	if IsRomanNumeral(operand1Str) {
		calculator.operand1 = romanToArabic(strings.TrimSpace(operand1Str))
		calculator.operand2 = romanToArabic(strings.TrimSpace(operand2Str))
		switch calculator.operator {
		case "+":
			result = calculator.Add()
		case "-":
			result = calculator.Subtract()
		case "*":
			result = calculator.Multiply()
		case "/":
			result = calculator.Divide()
		}
	} else {
		switch calculator.operator {
		case "+":
			result = calculator.Add()
		case "-":
			result = calculator.Subtract()
		case "*":
			result = calculator.Multiply()
		case "/":
			result = calculator.Divide()
		}
	}
	fmt.Printf("Результат: %d\n", result)
}
