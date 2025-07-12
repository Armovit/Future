package main

import "fmt"

func main() {
	var a string = "Юрий"
	var c int = 5
	var count int = 5
	fmt.Println(a)
	greetSum(c)
	c = c + 5
	greet(c)
	_ = count
	for i := 0; i < count; i++ {
		c += c
		greet(c)
	}

}

func greet(a int) {
	fmt.Println(a)
}

func greetSum(a int) {
	a = a + 10
	fmt.Println(a)
}

func greetNow(name string) {

}
