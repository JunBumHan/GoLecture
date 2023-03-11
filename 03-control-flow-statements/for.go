package main

import "fmt"

func main() {

	var i int = 5

	// for문
	for i = 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 초기문 생략
	i = 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	// 후처리 생략
	for i = 0; i < 10; {
		fmt.Println(i)
		i++
	}

	// 조건문, 후처리 생략
	i = 0
	for i < 10 { // == for ; i < 10; {}
		fmt.Println(i)
		i++
	}

	// 무한 반복문
	i = 0
	for {
		fmt.Println(i)
		if i == 10 {
			break
		}
		i++
	}

	// 레이블
OutLable:
	for i = 0; i < 10; i++ {
		for j := 0; j < 10; i++ {
			fmt.Println(i, j)
			if i == 5 {
				break OutLable // 모든 for문을 나감
			}
			if j == 3 {
				break // break가 속한 for문 (초기문이 j := 0이 for문을 나감)
			}
		}
	}

}
