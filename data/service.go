package data

type Service interface {
	UbahStatus()
	CekAvailability() bool
}