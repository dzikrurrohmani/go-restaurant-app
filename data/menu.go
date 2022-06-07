package data

type Menu struct {
	Name string
	Price int
	Stock int
}

func (m *Menu) UbahStatus() {
	m.Stock = m.Stock-1
}

func (m Menu) CekAvailability() bool {
	return m.Stock != 0
}