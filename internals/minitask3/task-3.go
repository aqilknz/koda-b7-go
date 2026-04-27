package minitask3

import (
	"fmt"
	"slices"
)

func InsertDataToSlice() {
	data := []int{50, 75, 66, 20, 32, 90}

	data = slices.Insert(data, 3, 88)

	fmt.Println(data)

	for i, v := range data {
		fmt.Printf("Index %d: %d\n", i, v)
	}
}
