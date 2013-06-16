package gounqlite_test

import (
	"fmt"
	"gounqlite"
)

func ExampleThreadsafe() {
	if r := gounqlite.Threadsafe(); r {
		fmt.Println(r)
		// Output: true
	} else {
		fmt.Println(r)
		// Output: false
	}
}
