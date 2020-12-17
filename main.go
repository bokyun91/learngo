package main

import (
	"fmt"
	"strings"
)

// 자바와 다르게 변수명 뒤에 type 작성. return 타입은 arg 뒤에 작성
func multiply(a int, b int) int {
	return a * b
}

// 여러 개 반환 가능
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// 같은 타입의 arg를 여러 개 받을 경우 ...으로 받을 수 있음
func repeatMe(words ...string) {
	fmt.Println(words)
}

// 대문자로 시작할 경우(함수, 변수) -> public
func main() {
	// 상수 const, 변수 var
	// 변수 선언 => var/const + 변수명 + 자료형 or 변수명 := 값(자료형은 자동, 단 func 밖에서는 안됨)
	var name string = "nico"
	name = "lyn"
	fmt.Println(name)
	fmt.Println(multiply(2, 2))

	// 여러 변수 return 받을 때. 하나만 사용할 경우 _ .
	totalLength, upperName := lenAndUpper("nico")
	totalLength2, _ := lenAndUpper("pbk")
	fmt.Println(totalLength, upperName)
	fmt.Println(totalLength2)

	repeatMe("nico", "lynn", "dal", "marl")
}
