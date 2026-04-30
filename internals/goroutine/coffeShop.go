package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func barista(pesananKopi string, idBarista int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Barista %d=> lagi buat: %s...\n", idBarista, pesananKopi)
	time.Sleep(2 * time.Second)
	fmt.Printf("Barista %d=> telah selesai membuat kopi: %s\n", idBarista, pesananKopi)
}

func CoffeeShop() {
	var wg sync.WaitGroup
	daftarPesanan := []string{"Americano", "Latte", "Matcha"}

	for i, kopi := range daftarPesanan {
		wg.Add(1)
		id := i + 1
		// wg.Go(func() {
		// 	barista(kopi, id)
		// })
		go barista(kopi, id, &wg)

	}
	wg.Wait()
	fmt.Println("semua pesanan sudah")
	fmt.Println("Toko tutup")

}
