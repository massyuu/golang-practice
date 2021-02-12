package main

import "fmt"

/*
ポインタの動作確認
*/
func main() {

	var a1 int
	var b1 int
	var p1 *int

	a1 =123
	p1 = &a1
	b1 = a1
	fmt.Println(a1)
	fmt.Println(p1)
	fmt.Println(b1)
	*p1 = 456
	fmt.Println(a1)
	fmt.Println(p1)
	fmt.Println(b1)
}
