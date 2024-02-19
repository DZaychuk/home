package main

import (
	"fmt"
	"os"
)

func main() {
	var size int
	fmt.Fscan(os.Stdin, &size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
