package entitas

type Cart struct {
	UserID   int       `json:"userId"`
	Products []Product `json:"products"`
}

type CartRequest struct {
	KodeProduk string `json:"kodeProduk"`
	Kuantitas  int    `json:"kuantitas"`
}
