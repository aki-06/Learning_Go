package main

import "fmt"

// 二つ以上の引数が同じ型の場合は、
// 最後の型を残して省略できる
func add(x, y int) int {
	return x + y
}

func main(){
	fmt.Println(add(12, 7))
}