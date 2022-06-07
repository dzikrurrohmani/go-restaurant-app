package main

import (
	"bufio"
	"flag"
	"fmt"
	"live-code-3-1/data"
	"live-code-3-1/procedure"
	"os"
	"strings"
	"time"
)

func main() {
	/*
		Applikasi warung makan

		Warung Makan Bahari (WMB) adalah sebuah warteg yg sangat laku dan terkenal di daerah ragunan.
		Karena terkenal hingga sering direview oleh youtuber, WMB mendapatkan suntikan dana dari investor.
		Setelah dana investasi cair, WMB melakukan upgrade terhadap tempat makan nya yaitu meja makan dan sistem pemesanan dan pembayaran.
		Kini meja makan sudah terpisah-pisah dan memiliki nomor masing-masing.
		Pemesanan makan pun sudah menggunakan aplikasi di kasir, dan pembayaran dilakukan setelah selesai makan.
		ketentuan :
		-total meja makan ada 30 meja
		-sebelum ke meja makan, konsumen harus pesan makanan di kasir, baru akan diarahkan ke meja tertentu
		-untuk meja yg sudah ada konsumen, tidak bisa lagi dipilih
		-saat akan memilih meja, kasir dapat melihat meja mana yg masih kosong / available
		-Kasir Dapat melihat daftar makanan/mencari makanan berdasarkan kode/nama
		-saat konsumen memesan makanan, kasir menginput data pesanan, nama konsumen, dan juga nomor meja yg akan ditempati konsumen
		-setelah konsumen selesai makan, dapat melakukan pembayaran di kasir dan kemudian meja tersebut sudah available untuk konsumen berikutnya

		repo: livecode-3-wmb
	*/
	var state, tindakan int
	menu, meja := DataHariIni(configWithFlag())
	pesanan := []data.Pesanan{}
	// fmt.Println(len(meja))
	// fmt.Println(len(menu))
	// for _, structNya := range menu {
	// 	fmt.Println(structNya.(*data.Menu).Name)
	// }

	fmt.Println("Hai Kasir")
	fmt.Println("Selamat Bekerja di Warung Bahari")
	fmt.Printf("Ada %d Jenis Menu dan %d Meja yang bisa dipesan!\nAplikasi terbuka dalam :", len(menu), len(meja))
	for delay := 0; delay < 6; delay++ {
		fmt.Printf("\r\t\t\t%2d", 5-delay)
		time.Sleep(time.Second / 5)
	}
	// fmt.Printf("\033[%d;%dH", line, col) // Set cursor position

	state = 0
	for {
		if state == 0 {
			fmt.Print("\033[2J") //Clear screen
			fmt.Print("\n", strings.Repeat("-", 55))
			fmt.Println("\n- Sistem Pemesanan dan Pembayaran Warung Makan Bahari -")
			fmt.Println(strings.Repeat("-", 55))

			fmt.Println("Silakan pilih tindakan:", strings.Repeat(" ", 13), "Meja Tersedia:", procedure.MejaAvail(meja))

			fmt.Println("1. Daftar Menu")
			fmt.Println("2. Buat Pesanan")
			fmt.Println("3. Daftar Pesanan")
			fmt.Print("4. Keluar\n")
			fmt.Print("Masukkan angka: ")
			fmt.Scanln(&tindakan)
			if tindakan == 1 {
				state = 1
			}
			if tindakan == 2 {
				state = 2
			}
			if tindakan == 3 {
				state = 3
			}
			if tindakan == 4 {
				break
			}
			tindakan = 0
		} else if state == 1 {
			fmt.Print("\033[2J") //Clear screen
			procedure.PrintMenuSemua(menu)
			validInput := bufio.NewScanner(os.Stdin)
			fmt.Print("\nketik apapun untuk kembali ke menu utama -> ")
			validInput.Scan()
			state = 0

		} else if state == 2 {
			pesanState := 0
			pesanTemp := data.Pesanan{}
			for {
				if pesanState == 0 {
					var mejaInput, mejaPilihan int
					fmt.Print("\033[2J") //Clear screen
					fmt.Print("Butuh berapa meja : ")
					fmt.Scanln(&mejaInput)
					if mejaInput < procedure.MejaAvail(meja) {
						procedure.PrintMejaSemua(meja)
						i := 0
						for {
							fmt.Printf("Meja ke %d : ", i+1)
							fmt.Scanln(&mejaPilihan)
							if mejaPilihan < len(meja)+1 && meja[mejaPilihan-1].CekAvailability() {
								pesanTemp.MejaPesanan = append(pesanTemp.MejaPesanan, &meja[mejaPilihan-1])
								i++
							} else {
								fmt.Println("Meja tidak tersedia.")
							}
							if i == mejaInput {
								break
							}
						}
						pesanState = 1
					} else {
						fmt.Println("Jumlah meja tersedia tidak mencukupi.")
					}
				} else if pesanState == 1 {
					var menuPilihan, menuJumlah int
					fmt.Print("\033[2J") //Clear screen
					procedure.PrintMenuSebagian(menu)
					for {
						fmt.Print("Masukkan nomor dari menu yang dipilih : ")
						fmt.Scanln(&menuPilihan)
						if menuPilihan < len(menu)+1 && menu[menuPilihan-1].CekAvailability() {
							fmt.Printf("Menu %s tersedia %d item, ingin pesan berapa : ", menu[menuPilihan-1].(*data.Menu).Name, menu[menuPilihan-1].(*data.Menu).Stock)
							fmt.Scanln(&menuJumlah)
							menuPesanTemp := data.Menu{
								Name:  menu[menuPilihan-1].(*data.Menu).Name,
								Price: menu[menuPilihan-1].(*data.Menu).Price,
								Stock: menuJumlah,
							}
							pesanTemp.MenuPesanan = append(pesanTemp.MenuPesanan, menuPesanTemp)

							validInput := bufio.NewScanner(os.Stdin)
							fmt.Print("Apakah ingin memesan menu lain?\n(ketik y jika iya atau n jika tidak)\n-> ")
							validInput.Scan()
							if validInput.Text() != "y" {
								pesanState = 2
								break
							}
						} else {
							fmt.Println("Menu tidak tersedia.")
						}
					}
				} else if pesanState == 2 {
					fmt.Println("Berikut informasi pesanan anda:")
					pesanTemp.HitungBiaya()
					fmt.Printf("%T", pesanTemp.MejaPesanan[1])
					pesanTemp.PrintPesanan()
					validInput := bufio.NewScanner(os.Stdin)
					fmt.Print("Apakah pesanan tersebut sudah benar?\n(masukkan y jika benar atau n jika salah)\n-> ")
					validInput.Scan()
					if validInput.Text() == "y" {
						namaPemesan := bufio.NewScanner(os.Stdin)
						fmt.Print("Ingin membuat pesanan ini atas nama siapa : ")
						namaPemesan.Scan()
						pesanTemp.Nama = namaPemesan.Text()
						for _, mejaNya := range pesanTemp.MejaPesanan {
							(*mejaNya).UbahStatus()
						}
						for _, menuPil := range pesanTemp.MenuPesanan {
							for _, menuNya := range menu {
								if menuNya.(*data.Menu).Name == menuPil.Name {
									for i := 0; i < menuPil.Stock; i++ {
										menuNya.UbahStatus()
									}
								}
							}
						}
						pesanan = append(pesanan, pesanTemp)
						fmt.Printf("Pesanan atas nama %s berhasil dibuat.", pesanan[len(pesanan)-1].Nama)
						validInput := bufio.NewScanner(os.Stdin)
						fmt.Print("\nketik apapun untuk kembali ke menu utama -> ")
						validInput.Scan()
						state = 0
						break
					} else {
						pesanState = 0
						continue
					}
				}
			}
		}
	}
}

func configWithFlag() Config {
	conMenu := flag.String("menu", "menu.txt", "Menu Hari Ini")
	conMeja := flag.Int("meja", 30, "Meja Hari Ini")

	flag.Parse()
	return Config{
		Menu: *conMenu,
		Meja: *conMeja,
	}
}
