package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string = "koda"
	var b int
	var c float32
	const d int8 = 21
	fmt.Printf("%s, %d, %.3f\n", a, b, c)
	fmt.Println(d)
	fmt.Print("Hai ")
	fmt.Println("Hello World!")

	// konversi
	isSame := b == int(c)
	fmt.Println(isSame)
	kata := a + strconv.Itoa(int(d))
	fmt.Println(kata)

	multiple(2, 7)
	total, rerata := multipleApp(67, 89, 85, 90)
	fmt.Printf("Sum = %d\nAvg = %.2f", total, rerata)

}

// multiple input dengan 1 output
func multiple(angka1 int, angka2 int) int {
	return angka1 * angka2
}
func multipleApp(mtk uint, eng uint, indo uint, ipa uint) (sum uint, avg float32) {
	sum = mtk + eng + indo + ipa
	avg = float32(sum) / float32(4)
	return sum, avg
}
