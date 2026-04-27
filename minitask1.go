package main

import (
	"fmt"
	"math"
)

func main() {
	jariJari := 21.0
	// k := kelilingLingkaran(jariJari)
	// fmt.Printf("Keliling : %.2f\n", k)
	// l := luasLingkaran(jariJari)
	// fmt.Printf("Luas : %.2f\n", l)

	luas, keliling := HitungLingkaran(jariJari)

	fmt.Printf("Luas     : %.2f\n", luas)
	fmt.Printf("Keliling : %.2f\n", keliling)
}

func HitungLingkaran(r float64) (float64, float64) {
	return luasLingkaran(r), kelilingLingkaran(r)
}

func kelilingLingkaran(r float64) float64 {
	return 2 * math.Pi * r
}

func luasLingkaran(r float64) float64 {
	return math.Pi * r * r
}
