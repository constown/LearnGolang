package main

import (
	"fmt"
)

func main() {
	fmt.Println("1、==========")
	for i := 1; i <= 9; i++ {
		for n := 1; n <= i; n++ {
			fmt.Printf("%d*%d=%d\t", n, i, i*n)
		}
		fmt.Println()
	}

	fmt.Println("2、==========")
	for i := 9; i >= 1; i-- {
		for n := 9; n >= i; n-- {
			fmt.Printf("%d*%d=%d\t", n, i, i*n)
		}
		fmt.Println()
	}

	fmt.Println("3、==========")
	for i := 9; i >= 1; i-- {
		for n := 1; n <= i; n++ {
			fmt.Printf("%d*%d=%d\t", n, i, i*n)
		}
		fmt.Println()
	}

	fmt.Println("4、==========")
	for i := 1; i <= 9; i++ {
		for n := 9; n >= i; n-- {
			fmt.Printf("%d*%d=%d\t", i, n, i*n)
		}
		fmt.Println()
	}

	fmt.Println("5、==========")
	m, l := 9, 9
	for m >= 1 {
		if l >= 1 {
			fmt.Printf("%d*%d=%d\t", m, l, m*l)
			l--
		} else {
			m--
			l = m
			fmt.Println()
		}
	}

}
