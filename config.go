package main

type Config struct {
	Menu string
	Meja int
}

// func DataHariIni(c Config) ([]data.Menu, []data.Meja) {
// 	panic("")
// }

func DataHariIni(c Config) (string, int) {
	return c.Menu, c.Meja
}
