package main

import "fmt"

func main(){
	for i := 1; i <= 20; i++ {
		if i % 2 == 0{
			fmt.Println("bilangan genap: ", i)
		}
	}

	total := 0

	for number := 1; number <= 100; number++{
		total = total + number
	}

	println("jumlah semua adalah", total)
}

