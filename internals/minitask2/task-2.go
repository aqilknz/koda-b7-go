package minitask2

import "fmt"

func Segitiga(n int) {
	for i := range n {
		for range i + 1 {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
