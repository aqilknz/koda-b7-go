package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func mandi(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Mulai: persiapan mandi...")
	time.Sleep(3 * time.Second)
	fmt.Println("Selesai: sudah mandi")
}
func buatKopi(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Mulai: panaskan air panas ...")
	time.Sleep(3 * time.Second)
	fmt.Println("Selesai: kopi sudah siap")

}
func sarapan(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Mulai: siapkan bahan-bahan...")
	time.Sleep(4 * time.Second)
	fmt.Println("Selesai: makanan sudah siap")

}
func kamarRapi(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Mulai: bersikan kasur pakai sapu...")
	time.Sleep(3 * time.Second)
	fmt.Println("Selesai: kamar telah sudah rapi")

}

func Activity() {
	var wg sync.WaitGroup

	wg.Add(4)
	fmt.Println("Alarm Bangun Pagi!!")
	go mandi(&wg)
	go buatKopi(&wg)
	go sarapan(&wg)
	go kamarRapi(&wg)

	// fmt.Println("sedang melakukan aktivitas...")

	wg.Wait()
	fmt.Println("Berangkat Kerja")

}
