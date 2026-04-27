package minitask4

import "fmt"

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

func ShowBiodata() {

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
