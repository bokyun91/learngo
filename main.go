package main

import "fmt"

// 대문자로 시작할 경우(함수, 변수) -> public
func main() {
	// 상수 const, 변수 var
	// 변수 선언 => var/const + 변수명 + 자료형
	var name string = "nico"
	name = "lyn"
	fmt.Println(name)
}
