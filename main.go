package main

import (
	"golang.org/x/tour/tree"
	"math/rand"
	"time"
)
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	RecursiveWalk(t, ch)
	close(ch)
}
func RecursiveWalk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		RecursiveWalk(t.Left, ch)
	}
	if t.Right != nil {
		RecursiveWalk(t.Right, ch)
	} else {
		return
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	var v1, v2 int
	var ok1, ok2 bool
	for {
		v1, ok1 = <-ch1
		v2, ok2 = <-ch2
		fmt.Printf("%d ", v1)
		fmt.Printf("%d ", v2)
		fmt.Printf("%t ", ok1)
		fmt.Printf("%t ", ok2)
		if v1 != v2 || !ok1 || !ok2 {
			return false
		}
		fmt.Println()
	}
	return true
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("RESULT: %t", Same(tree.New(1), tree.New(1)))
	fmt.Print("\nExit")
}
