package main

import (
	"fmt"
	"os"
)

func IndexwithValue() {
	for i, arg := range os.Args {
		fmt.Printf("Index: %d Value: %s\n", i, arg)
	}
}
