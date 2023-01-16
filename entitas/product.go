package entitas

type Product struct {
	KodeProduk string `json:"kodeProduk"`
	NamaProduk string `json:"namaProduk"`
	Kuantitas  int    `json:"kuantitas"`
}
