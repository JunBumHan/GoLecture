package main

import "fmt"

func main() {
	a, b := 5, 10

	if a > b {
		fmt.Println("a > b true")
	} else {
		fmt.Println("a < b ture")
	}

	// 쇼트서킷 조심!
	if a > b && a < b {
		fmt.Println("쇼트서킷 조심!")
	} else if a > b || a < b {
		fmt.Println("쇼트서킷 조심!")
	}
}
