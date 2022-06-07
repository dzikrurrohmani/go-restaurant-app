package data

type Meja struct {
	Nomor int
	Availability bool
}

func (m *Meja) UbahStatus() {
	m.Availability = !m.Availability
}

func (m Meja) CekAvailability() bool {
	return m.Availability
}