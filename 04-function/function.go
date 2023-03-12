package main

import "fmt"

func Add(value1, value2 int) int { // == func add(value1, value2 int) int {}
	return value1 + value2
}

// 여러개 반환하기
func divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}

	return a / b, true
}

// 변수명을 지정해 반환하기
func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}
	result = a / b
	success = true
	return
}

// defer
func Hello() {
	defer fmt.Println("A")
	defer fmt.Println("B")
	fmt.Println("A!")
}

// 함수 타입 변수
func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func getOperator(op string) func(int, int) int {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}

}

// 함수 리터럴 2
func writeHello(writer func(string)) {
	writer("Hello world")
}

func main() {

	// 함수 리터럴 1
	i := 0
	f := func() {
		i += 10
	}

	f()

	writeHello(func(msg string) {
		fmt.Printf(msg)
	})
}
