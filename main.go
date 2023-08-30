package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

/*
The program initializes a "heap" with 10 objects.
Every iteration, the mark phase marks random objects as Alive = false.
The sweep phase prints the "collected" objects and "recreates" them by setting Alive = true.
The heap's status is printed with colors: green for alive objects and red for dead ones.
This is a vastly simplified model of a garbage collector.
*/

type Object struct {
	ID    int
	Alive bool
}

func main() {
	rand.Seed(time.Now().UnixNano())
	heap := createObjects(10)

	for {
		color.Cyan("Beginning mark phase...")
		mark(heap)

		color.Cyan("Beginning sweep phase...")
		sweep(heap)

		printHeap(heap)
		time.Sleep(3 * time.Second)
	}
}

func createObjects(n int) []*Object {
	objects := make([]*Object, n)
	for i := 0; i < n; i++ {
		objects[i] = &Object{ID: i, Alive: true}
	}
	return objects
}

func mark(heap []*Object) {
	// Randomly mark some objects as not alive
	for _, obj := range heap {
		if rand.Intn(2) == 0 {
			obj.Alive = false
		}
	}
}

func sweep(heap []*Object) {
	for _, obj := range heap {
		if !obj.Alive {
			color.Red("Object %d collected.", obj.ID)
			obj.Alive = true // "recreate" the object for the sake of this demonstration
		}
	}
}

func printHeap(heap []*Object) {
	fmt.Println("Heap status:")
	for _, obj := range heap {
		if obj.Alive {
			color.Green("Object %d: Alive", obj.ID)
		} else {
			color.Red("Object %d: Dead", obj.ID)
		}
	}
}
