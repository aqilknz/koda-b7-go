package goroutine

import (
	"fmt"
	"sync"
)

type Pesan struct {
	Pengirim string
	Isi      string
}

func pengirimPesan(pengirim string, isi string, ch chan<- Pesan, wg *sync.WaitGroup) {
	defer wg.Done()
	pesan := Pesan{
		Pengirim: pengirim,
		Isi:      isi,
	}
	ch <- pesan

}

func papanTulis(ch <-chan Pesan, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Papan Tulis Keluarga")
	for pesan := range ch {
		fmt.Printf("Pengirim %s : %s\n", pesan.Pengirim, pesan.Isi)
	}

}

func MessegeBoard() {
	papanChn := make(chan Pesan)
	var wgPapan sync.WaitGroup
	var wgPengirim sync.WaitGroup
	// go papanTulis(papanChn, &wg)

	wgPapan.Add(1)
	go papanTulis(papanChn, &wgPapan)

	wgPengirim.Add(1)
	go pengirimPesan("Aku", "Jangan makan makanan ku", papanChn, &wgPengirim)
	wgPengirim.Add(1)
	go pengirimPesan("Ibu", "Kasur dibersihkan yak", papanChn, &wgPengirim)

	wgPengirim.Wait()
	close(papanChn)
	wgPapan.Wait()
	fmt.Println("Program Selesai")

}
