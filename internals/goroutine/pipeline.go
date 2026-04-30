package goroutine

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

// Tahap 1: Generator
func genAngka(n int, hasil chan<- int) {
	for i := 1; i <= n; i++ {
		hasil <- i
	}
	close(hasil)
}
func filterGenap(angka <-chan int, hasil chan<- int) {
	for num := range angka {
		if num%2 == 0 {
			hasil <- num
		}
	}
	close(hasil)
}

func kuadratAngka(angka <-chan int, hasil chan<- int) {
	for num := range angka {
		hasil <- num * num
	}
	close(hasil)
}

func Pipeline() {
	var wg sync.WaitGroup
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan angka (N): ")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	n, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error: Harap masukkan angka yang valid!")
		return
	}

	chGenerator := make(chan int)
	chFilter := make(chan int)
	chHasil := make(chan int)

	wg.Add(3)

	go func() {
		defer wg.Done()
		genAngka(n, chGenerator)
	}()

	go func() {
		defer wg.Done()
		filterGenap(chGenerator, chFilter)
	}()

	go func() {
		defer wg.Done()
		kuadratAngka(chFilter, chHasil)
	}()

	var kumpulanHasil []int

	for hasil := range chHasil {
		kumpulanHasil = append(kumpulanHasil, hasil)
	}

	fmt.Printf("Hasil Akhir: %v\n", kumpulanHasil)

	wg.Wait()
	fmt.Println("Program selesai")
}
