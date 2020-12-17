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

// 같은 타입의 arg여러 개 받을 경우 ...으로 받을 수 있음
func repeatMe(words ...string) {
	fmt.Println(words)
}

// naked return
func lenAndUpper2(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// function 끝나고 또 다른 것 실행 시킴 defer -> return 후 실행
func lenAndUpper3(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func superAdd(numbers ...int) int {
	total := 0
	// range : array에 Loop 사용할 수 있도록 해줌. index
	for _, number := range numbers {
		total += number
	}

	// 평소 쓰던 방식
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }

	return total

}

func canIDrink(age int) bool {
	if age < 18 {
		return false
	}
	return true
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

	fmt.Println(lenAndUpper2("pbk"))
	fmt.Println(lenAndUpper2("pbk"))

	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)

	fmt.Println(canIDrink(16))
}
