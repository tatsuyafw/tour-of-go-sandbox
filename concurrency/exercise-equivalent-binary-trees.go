package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func WalkImpl(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	WalkImpl(t.Left, ch)
	ch <- t.Value
	WalkImpl(t.Right, ch)
}

// Walk wals the tree t sending all values
// from the tree to the cannel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkImpl(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if v1 != v2 {
			return false
		}
		if !ok1 && !ok2 {
			break
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for value := range ch {
		fmt.Printf("value = %v\n", value)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
