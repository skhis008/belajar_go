package main

import "fmt"

type Film struct {
	judul string
	tahun int
	rating float64
}

func main(){
	daftarFilm := []Film{
		{judul: "the incredibles", tahun: 2004, rating: 8.8},
		{judul: "the dark knight", tahun: 2008, rating: 9.0},
		{judul: "Annunaki", tahun: 2015, rating: 10.0},
	}

	for i, film := range daftarFilm{
		fmt.Printf("%d. %s (%d) - Rating: %.1f\n", i+1, film.judul, film.tahun, film.rating)
	}
}
