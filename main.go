package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	// minitask-1
	radius := 21
	area, circumference := HitungLingkaran(radius)
	fmt.Printf("Luas     : %.2f\n", area)
	fmt.Printf("Keliling : %.2f\n", circumference)

	// minitask-2
	segitiga(5)

	// minitask-3
	insertDataToSlice()

	// minitask-4
	type Pendidikan struct {
		nama    string
		jurusan string
	}

	type Biodata struct {
		name      string
		foto      string
		email     string
		umur      int
		telp      string
		isMerried bool
		education []Pendidikan
	}

	aqil := Biodata{
		name:      "Ahmad Aqil",
		foto:      "profile.jpg",
		email:     "khairunnz123@gmail.com",
		umur:      21,
		telp:      "08181812323",
		isMerried: false,
		education: []Pendidikan{
			{
				nama:    "SMK N 1 Adiwerna",
				jurusan: "TKJ",
			},
			{
				nama:    "Universitas Siber Asia",
				jurusan: "Informatika",
			},
		},
	}
	fmt.Println(aqil)

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

func segitiga(n int) {
	for i := range n {
		for range i + 1 {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func insertDataToSlice() {
	data := []int{50, 75, 66, 20, 32, 90}

	data = slices.Insert(data, 3, 88)

	fmt.Println(data)

	for i, v := range data {
		fmt.Printf("Index %d: %d\n", i, v)
	}
}
