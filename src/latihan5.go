package main

import "fmt"

func kali(a, b int) int{
	return a * b
}

func sapaWaktu(nama, waktu string) string{
	fmt.Println("selamat", waktu, nama)
	return "selamat " + waktu + " " + nama
}

func luasPersegiPanjang(panjang, lebar int) int{
	return panjang * lebar
}

func main() {
	hasil := kali(5, 10)
	fmt.Println("hasil perkalian adalah", hasil)
	sapaWaktu("Alice", "pagi")
	luas := luasPersegiPanjang(5, 10)
	fmt.Println("luas persegi panjang adalah", luas)
}

