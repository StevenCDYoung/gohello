package extral

import "fmt"

func Read() {
	fmt.Println("I'm reading...")
}

func Write() {
	fmt.Println("I'm writing...")
}

//This function will call the another function from b.go which is the same package
func Callrun() {
	run()
}

//This function can only be called by others that is in the same package (extral)
func fly() {
	fmt.Println("I'm flying...")
}
