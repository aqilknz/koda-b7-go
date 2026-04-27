package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aqilknz/koda-b7-go/internals/minitask1"
	"github.com/aqilknz/koda-b7-go/internals/minitask2"
	"github.com/aqilknz/koda-b7-go/internals/minitask3"
	"github.com/aqilknz/koda-b7-go/internals/minitask4"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("===Minitask===")
		fmt.Println("1. Hitung Luas & Keliling Lingkaran")
		fmt.Println("2. Segitiga Siku-Siku")
		fmt.Println("3. Manipulasi Slice")
		fmt.Println("4. Biodata")
		fmt.Println("0. Keluar")
		fmt.Println("")
		fmt.Print("Pilih Menu: ")

		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		choose, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Error: Masukkan angka yang valid!")
			continue
		}

		switch choose {
		case 1:
			fmt.Print("Masukkan jari-jari: ")
			scanner.Scan()
			rStr := scanner.Text()
			r, _ := strconv.Atoi(rStr)

			area, circumference := minitask1.HitungLingkaran(r)
			fmt.Printf("Hasil => Luas: %.2f dan Keliling: %.2f\n", area, circumference)

		case 2:
			fmt.Print("Masukkan tinggi segitiga: ")
			scanner.Scan()
			tStr := scanner.Text()
			t, _ := strconv.Atoi(tStr)
			minitask2.Segitiga(t)

		case 3:
			minitask3.InsertDataToSlice()

		case 4:
			minitask4.ShowBiodata()
		case 0:
			fmt.Println("Thanks")
			return

		default:
			fmt.Println("Pilihan tidak tersedia")
		}

	}

}
