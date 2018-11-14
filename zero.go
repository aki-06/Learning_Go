package main

import "fmt"

/*
変数に初期値を与えずに宣言すると、ゼロ値( zero value )が与えられます。

数値型(int,floatなど): 0
bool型: false
string型: "" (空文字列( empty string ))
*/

func main(){
	var i int  // 0
	var f float64  // 0
	var b bool  // false
	var s string  // (空文字)
	fmt.Println(i, f, b, s)
}