package procedure

import (
	"fmt"
	"live-code-3-1/data"
)

func PrintMenuSemua(sliceMenu []data.Service) {
	fmt.Println("----------------PILIHAN MENU----------------")
	fmt.Println("\n| No |       Nama       |  Price  |  Stok  |")
	for ke, menuNya := range sliceMenu {
		fmt.Printf("| %-2d |%-18s|  %5d  |   %2d   |\n", ke+1, menuNya.(*data.Menu).Name, menuNya.(*data.Menu).Price, menuNya.(*data.Menu).Stock)
	}
}

func PrintMenuSebagian(sliceMenu []data.Service) {
	fmt.Println("Menu yang tersedia:")
	for ke, menuNya := range sliceMenu {
		if menuNya.CekAvailability() {
			fmt.Printf("%-2d : %-18s\n", ke+1, menuNya.(*data.Menu).Name)
		}
	}
}
