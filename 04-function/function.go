package main

func add(value1, value2 int) int { // == func add(value1, value2 int) int {}
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

func main() {

}
