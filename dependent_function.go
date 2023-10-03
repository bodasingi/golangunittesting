// main.go
package main

func FunctionB(input int) int {
	return input * 2
}

func FunctionA(input int) int {
	result := FunctionB(input)
	return result + 1
}
