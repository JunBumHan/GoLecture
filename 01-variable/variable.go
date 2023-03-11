package main

// a := 5 (불가능)
const ABC = 5

func main() {

	// var 키워드를 이용해 선언
	var var1 int

	// := 키워드를 이용해 선언한다. (데이터 타입을 추론한다)
	var2 := 5 // int형으로 자동적으로 타입 추론 된다
	// 여러개를 묶어 선언하기
	var3, var4, var5 := 0, 1, 2
	var var6, var7, var8 int

	// 괄호로 묶어 선언하기
	var (
		a int
		b int
		c int
	)

}
