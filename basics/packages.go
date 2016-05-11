package main

import(
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(100)
	fmt.Println("My favorite number is", rand.Intn(10))
}
