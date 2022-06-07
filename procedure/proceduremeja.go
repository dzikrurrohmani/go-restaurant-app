package procedure

import (
	"fmt"
	"live-code-3-1/data"
)

func MejaAvail(sliceMeja []data.Service) int {
	jumlahAvail := 0
	for _, mejaNya := range sliceMeja {
		if mejaNya.CekAvailability() {
			jumlahAvail++
		}
	}
	return jumlahAvail
}

func PrintMejaSemua(sliceMeja []data.Service) {
	var mejaAvailText string
	fmt.Println("Ketersediaan meja:")
	for _, mejaNya := range sliceMeja {
		if mejaNya.CekAvailability() {
			mejaAvailText = "Tersedia"
		} else {
			mejaAvailText = "Tidak tersedia"
		}
		fmt.Printf("Meja ke - %d : %v\n", mejaNya.(*data.Meja).Nomor, mejaAvailText)
	}
}
