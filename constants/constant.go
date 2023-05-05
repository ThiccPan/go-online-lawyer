package constants

const JWT_SECRET = "secret"

type KonsultasiStatus string

const (
	PEMBAYARAN KonsultasiStatus = "PEMBAYARAN"
	DIPROSES KonsultasiStatus = "DIPROSES"
	SELESAI KonsultasiStatus = "SELESAI"
	BATAL KonsultasiStatus = "DIBATALKAN"
)