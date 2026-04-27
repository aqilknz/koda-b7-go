package main

import (
	"fmt"
	"math"
)

func main() {
	radius := 21
	// k := kelilingLingkaran(radius)
	// fmt.Printf("Keliling : %.2f\n", k)
	// l := luasLingkaran(radius)
	// fmt.Printf("Luas : %.2f\n", l)

	area, circumference := HitungLingkaran(radius)

	fmt.Printf("Luas     : %.2f\n", area)
	fmt.Printf("Keliling : %.2f\n", circumference)
}

func HitungLingkaran(r int) (float32, float32) {
	return luasLingkaran(r), kelilingLingkaran(r)
}

func kelilingLingkaran(r int) float32 {
	return 2 * math.Pi * float32(r)
}

func luasLingkaran(r int) float32 {
	return math.Pi * float32(r*r)
}
