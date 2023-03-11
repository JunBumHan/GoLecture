package main

import "fmt"

func main() {
	a := 3

	// switch 문
	switch a {
	case 3:
		fmt.Println("a is 3")
	case 2:
		fmt.Println("a is 2")
	case 1:
		fmt.Println("a is 3")
	default:
		fmt.Println("?")
	}

	switch a = getInteger(); a { // siwtch 초기문
	case 3:
		fmt.Println("a is 3")
	case 2:
		fmt.Println("a is 2")
	case 1:
		fmt.Println("a is 3")
	default:
		fmt.Println("?")
	}

	switch { // == switch true {}
	case a == 1:
		fmt.Println("a is 1")
	case a == 2:
		fmt.Println("a is 2")
	case a == 3:
		fmt.Println("a is 3")
	}

	switch a { // fallthrough
	case 1:
		fmt.Println("a is 1")
		fallthrough
	case 2:
		fmt.Println("a is 2")
	case 3:
		fmt.Println("a is 3")
		fallthrough
	default:
		fmt.Println("a is ?")
	}

}

func getInteger() int {
	return 5
}
