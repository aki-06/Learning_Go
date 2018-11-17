package main

import "fmt"

/*
定数はconstで宣言

以下が定数で宣言できる
- character
- string
- boolean
- numeric

定数は:=で宣言できない
*/

const Pi = 3.14

func main(){
	const world = "世界"
	fmt.Println("Hello", world)
	fmt.Println("Happy", Pi, "Day")

	const True = true
	fmt.Println(True)
}