package main

import "fmt"

func main() {
	a, b := 10, 5

	// 산술연산자
	fmt.Println(a+b, a-b, a*b, a/b, a%b)

	// 비교 연산자
	fmt.Println(a == b, a != b, a < b, a > b, a <= b, a >= b)

	// 논리 연산자
	fmt.Println(a == b && a != b, a < b || a > b, !(a > b))

}
