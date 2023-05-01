package storage

type PengacaraStorer interface {
	GetAll()
	GetByID()
	Insert()
	Update()
	Delete()
}