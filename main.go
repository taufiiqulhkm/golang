package main

import (
	"fmt"
	"os"
)

// struct untuk menyimpan data teman
type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// slice untuk menyimpan data teman-teman kelas
var daftarTeman = []Teman{
	{Nama: "Taufiiqulhakim jubair", Alamat: "pondok cipta", Pekerjaan: "belajar", Alasan: "ingin mengetahiui seberapa bisa mengikuti golang ini"},
	{Nama: "moch basir", Alamat: "mustika jaya", Pekerjaan: "mahasiswa", Alasan: "ingin memahami bahasa golang"},
	{Nama: "ardhiasnyah", Alamat: "cikarang", Pekerjaan: "pegaawai", Alasan: "ingin belajar"},
}

// fungsi untuk menampilkan data teman berdasarkan nomor absen
func tampilkanDataTeman(absen int) {
	if absen <= len(daftarTeman) {
		teman := daftarTeman[absen-1]
		fmt.Printf("Data teman dengan absen %d:\n", absen)
		fmt.Printf("Nama: %s\n", teman.Nama)
		fmt.Printf("Alamat: %s\n", teman.Alamat)
		fmt.Printf("Pekerjaan: %s\n", teman.Pekerjaan)
		fmt.Printf("Alasan memilih kelas Golang: %s\n", teman.Alasan)
	} else {
		fmt.Println("Data teman tidak ditemukan.")
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Masukkan nomor absen teman.")
		return
	}

	var absen int
	_, err := fmt.Sscan(os.Args[1], &absen)
	if err != nil {
		fmt.Println("Masukkan nomor absen dalam bentuk angka.")
		return
	}

	tampilkanDataTeman(absen)
}
