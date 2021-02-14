package main

import (
	// "/extral"
	"fmt"

	"example.com/extral"
)

func main() {

	fmt.Println("Hello world!")

	fmt.Println("宽度：", WIDTH, "高度：", HEIGHT)
	dataCon()
	dance()
	fmt.Println(str)

	extral.Read()
	extral.Write()
	extral.Callrun()
	extral.Sing()
	extral.Dance()
}
