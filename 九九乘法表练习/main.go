package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for n := 1; n < 10; n++ {
			if n > i {
				fmt.Println()
				break
			} else {
				fmt.Printf("%d*%d=%d  ", n, i, i*n)
			}
		}
	}
}
