package data

import "fmt"

type Pesanan struct {
	Nomor       int
	Nama        string
	MejaPesanan []*Service
	MenuPesanan []Menu
	Total       int
}

func (p *Pesanan) HitungBiaya() {
	for _, menuNya := range p.MenuPesanan {
		p.Total += menuNya.Price * menuNya.Stock
	}
}

func (p Pesanan) PrintPesanan() {
	fmt.Println("Total meja yang dipesan:", len(p.MejaPesanan))
	fmt.Println("Menu yang dipesan:", len(p.MejaPesanan))
	for _, menuNya := range p.MenuPesanan {
		fmt.Printf("%-18sx  %5d * %2d = %d\n", menuNya.Name, menuNya.Price, menuNya.Stock, menuNya.Price*menuNya.Stock)
	}
	fmt.Println("Total biaya pesanan:", p.Total)
}
