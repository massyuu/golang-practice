package main

import "fmt"

func main() {
	var message = "Hello"
	fmt.Println(message)
	message = "world"
	fmt.Println(message)
	var testPrSt prSt
	testPrSt.x = 1
	testPrSt.y = 2
	testPrSt.clsPr()

	var testChSt chSt
	testChSt.x = 10
	testChSt.y = 20
	testChSt.z = "test"
	testChSt.prSt.clsPr()
	testChSt.clsSt()

}

type prSt struct{
	x int
	y int
}

type chSt struct{
	z string
	prSt
}

func (p *prSt) clsPr() {
	fmt.Println("Parent", p.x, p.y)
}

func (p *chSt) clsSt() {
	fmt.Println("child", p.z, p.prSt.y)
}
