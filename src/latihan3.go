package main

import "fmt"

func main() {
	umur := 18
	punyaTiket := true

	if umur >= 18 && punyaTiket {
		fmt.Println("anda boleh masuk")
	} else {
		fmt.Println("anda belum boleh masuk")
	}

	hari := "sabtu"

	switch hari {
	case "sabtu", "minggu":
		fmt.Println("hari libur")
	default:
		fmt.Println("hari kerja")
	}

	nilai := 75

	if nilai >= 90{
		fmt.Println("Anda lulus dan nilai anda adalah A")
	} else if nilai >= 70 {
		fmt.Println("Anda lulus dan Nilai anda adalah B ")
	} else {
		fmt.Println("anda tidak lulus dan nilai anda adalah C")
	}
}
