package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aqilknz/koda-b7-go/internals/goroutine"
	"github.com/aqilknz/koda-b7-go/internals/minitask1"
	"github.com/aqilknz/koda-b7-go/internals/minitask2"
	"github.com/aqilknz/koda-b7-go/internals/minitask3"
	"github.com/aqilknz/koda-b7-go/internals/minitask4"
	"github.com/aqilknz/koda-b7-go/internals/minitask6"
	"github.com/aqilknz/koda-b7-go/internals/minitask7"
	"github.com/aqilknz/koda-b7-go/internals/minitask8"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	person := minitask7.NewPerson("Ahmad Aqil", "Acropolis C.10 No.29", "0822929292")

	for {
		fmt.Println("===Minitask===")
		fmt.Println("1. Hitung Luas & Keliling Lingkaran")
		fmt.Println("2. Segitiga Siku-Siku")
		fmt.Println("3. Manipulasi Slice")
		fmt.Println("4. Biodata")
		fmt.Println("5. Baca File")
		fmt.Println("6. Print Data Personal")
		fmt.Println("7. Sapa")
		fmt.Println("8. Ganti Nama")
		fmt.Println("9. Payment")
		fmt.Println("10. CoffeShop")
		fmt.Println("11. Activity")
		fmt.Println("12. Message")
		fmt.Println("13. Pipeline")
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

		case 5:
			fmt.Print("Masukkan path file: ")
			scanner.Scan()
			filePath := strings.TrimSpace(scanner.Text())

			err := minitask6.ReadFile(filePath)
			if err != nil {
				fmt.Println("Gagal membaca file: ", err)
			}

		case 6:
			person.Print()

		case 7:
			fmt.Println(person.Greet())

		case 8:
			fmt.Println("old:\n", person.Greet())
			fmt.Print("\nMasukkan Nama: ")
			scanner.Scan()
			inputNama := strings.TrimSpace(scanner.Text())
			person.ChangeName(inputNama)
			fmt.Println("\nnew:\n", person.Greet())

		case 9:
			prices := []int{5000, 10000, 25000}
			invalidPrices := []int{-5000}

			fiktif := &minitask8.Fiktif{}

			checkout := minitask8.NewCheckout(minitask8.Bank{})
			checkout.Process(prices)

			checkout = minitask8.NewCheckout(minitask8.Online{})
			checkout.Process(prices)

			checkout = minitask8.NewCheckout(fiktif)
			checkout.Process(prices)

			// error handling
			checkout.Process(invalidPrices)

			fmt.Println("Riwayat Pembayaran Fiktif:")
			fmt.Println(fiktif.History)
		case 10:
			goroutine.CoffeeShop()
		case 11:
			goroutine.Activity()
		case 12:
			goroutine.MessegeBoard()
		case 13:
			goroutine.Pipeline()
		case 0:
			fmt.Println("Thanks")
			return

		default:
			fmt.Println("Pilihan tidak tersedia")
		}

	}

}
