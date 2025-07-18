package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

type Operation func(int, int) int

func add(a, b int) int { return a + b }
func sub(a, b int) int { return a - b }
func mul(a, b int) int { return a * b }
func pow(a, b int) int { return int(math.Pow(float64(a), float64(b))) }
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func randomSlice(n int) []int {
	rand.Seed(time.Now().UnixNano())
	s := make([]int, n)
	for i := range s {
		s[i] = rand.Intn(100) + 1
	}
	return s
}

func mapSlice(s []int, f func(int) int) []int {
	res := make([]int, len(s))
	for i, v := range s {
		res[i] = f(v)
	}
	return res
}

func reduceSlice(s []int, f func(int, int) int, init int) int {
	res := init
	for _, v := range s {
		res = f(res, v)
	}
	return res
}

func writeToFile(filename string, data []int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, v := range data {
		fmt.Fprintf(file, "%d\n", v)
	}
	return nil
}

func main() {
	ops := []Operation{add, sub, mul, pow, gcd}
	names := []string{"add", "sub", "mul", "pow", "gcd"}
	a, b := rand.Intn(50)+1, rand.Intn(20)+1
	for i, op := range ops {
		res := op(a, b)
		fmt.Printf("%s(%d, %d) = %d\n", names[i], a, b, res)
	}

	slice := randomSlice(10)
	fmt.Printf("Random slice: %v\n", slice)
	squared := mapSlice(slice, func(x int) int { return x * x })
	fmt.Printf("Squared: %v\n", squared)
	sum := reduceSlice(slice, add, 0)
	fmt.Printf("Sum: %d\n", sum)
	max := reduceSlice(slice, func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}, slice[0])
	fmt.Printf("Max: %d\n", max)
	err := writeToFile("numbers.txt", slice)
	if err == nil {
		fmt.Println("Slice written to numbers.txt")
	}
}
