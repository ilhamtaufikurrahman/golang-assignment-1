package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1]

	fmt.Println("Data nomor absen", args)

	noAbsen, _ := strconv.Atoi(args)

	var students = []biodata{
		{nama: "Ilham", alamat: "Kota Bandung", pekerjaan: "Mahasiswa", alasan: "Golang best"},
		{nama: "Taufik", alamat: "Kabupaten Bandung", pekerjaan: "Mahasiswa", alasan: "Golang Bagus"},
		{nama: "Rahman", alamat: "Kota Garut", pekerjaan: "Mahasiswa", alasan: "Golang performanya bagus"},
		{nama: "Fuad", alamat: "Kabupaten Garut", pekerjaan: "Mahasiswa", alasan: "Golang sangat mudah"},
		{nama: "Solehudin", alamat: "Kota Jakarta", pekerjaan: "Mahasiswa", alasan: "Golang fleksibel"},
		{nama: "Fikri", alamat: "Kota Bogor", pekerjaan: "Mahasiswa", alasan: "Golang awesome"},
		{nama: "Fadillah", alamat: "Kabupaten Bogor", pekerjaan: "Mahasiswa", alasan: "Golang very easy"},
		{nama: "Siti", alamat: "Kota Depok", pekerjaan: "Mahasiswa", alasan: "Golang tidak rumit"},
		{nama: "Azizah", alamat: "Kota Tanggerang", pekerjaan: "Mahasiswa", alasan: "Golang mudah dipelajari"},
		{nama: "Riski", alamat: "Kota Bekasi", pekerjaan: "Mahasiswa", alasan: "Golang hemat memori"},
	}

	var hasil biodata = cariData(noAbsen, students)

	if noAbsen > len(students) {
		fmt.Printf("Nomor absen %d tidak ada, data hanya ada %d", noAbsen, len(students))
	} else {
		fmt.Println("Nama\t\t:", hasil.nama)
		fmt.Println("Alamat\t\t:", hasil.alamat)
		fmt.Println("Pekerjaan\t:", hasil.pekerjaan)
		fmt.Println("Alasan\t\t:", hasil.alasan)

	}

}

type biodata struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func cariData(noAbsen int, data []biodata) biodata {
	var hasil biodata
	for i, student := range data {
		if noAbsen == i+1 {
			hasil = student
		}
	}

	return hasil
}
