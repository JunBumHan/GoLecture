package main

import "fmt"

func main() {
	a, b := 5, 10

	if a > b {
		fmt.Println("a > b true")
	} else {
		fmt.Println("a < b ture")
	}

	// 조건식에는 무조권 bool타입만이 와야함!
	// if 1 {
	// 	fmt.Println("True")
	// }

	// if 초기문
	if a, b := getInteger(), getInteger(); a > b {
		fmt.Println("a > b")
	} else if a < b {
		fmt.Println("a < b")
	} else {
		fmt.Println("==")
	}

	// 쇼트서킷 조심!
	if a > b && a < b {
		fmt.Println("쇼트서킷 조심!")
	} else if a > b || a < b {
		fmt.Println("쇼트서킷 조심!")
	}
}

func getInteger() int {
	return 5
}
