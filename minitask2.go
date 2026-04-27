package main

import "fmt"

func main() {
	segitiga(5)
}

// func segitiga(n int) {
// 	for i := 1; i <= n; i++ {
// 		for j := 1; j <= i; j++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
// }

func segitiga(n int) {
	for i := range n {
		for range i + 1 {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
