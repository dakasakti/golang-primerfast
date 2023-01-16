package entitas

type Filter struct {
	UserId     int    `json:"userId"`
	NamaProduk string `json:"namaProduk"`
	Kuantitas  int    `json:"kuantitas"`
}
