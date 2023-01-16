package repositories

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/dakasakti/golang/primerfast/entitas"
)

func TestCart(t *testing.T) {
	dataProduct := []entitas.Product{
		{
			KodeProduk: "DM-001",
			NamaProduk: "Product 1",
			Kuantitas:  1,
		},
		{
			KodeProduk: "DM-002",
			NamaProduk: "Product 2",
			Kuantitas:  1,
		},
	}

	mockDBCart := &mockDBCart{
		dataCart: []entitas.Cart{
			{
				UserID:   1,
				Products: dataProduct,
			},
			{
				UserID:   2,
				Products: dataProduct,
			},
		},
	}

	repo := NewRepoCart(mockDBCart, &mockRepoProduct{dataProduct: dataProduct})

	req := entitas.CartRequest{
		KodeProduk: "DM-003",
		Kuantitas:  1,
	}

	t.Run("tambahCart", func(t *testing.T) {
		err := repo.TambahCart(1, req)
		if err != nil {
			t.Errorf("error: %v", err)
		}

		cart, err := repo.cariCart(1)
		if err != nil {
			t.Errorf("error while finding cart: %v", err)
		}

		if len(cart.Products) != 1 {
			t.Errorf("error while adding cart, expected: %d, got: %d", 1, len(cart.Products))
		}
	})

	t.Run("tambahCartProdukTidakAda", func(t *testing.T) {
		req.KodeProduk = "DM-003"

		err := repo.TambahCart(1, req)
		if err == nil {
			t.Errorf("error should be return when adding a product that doesn't exist")
		}
	})

	t.Run("tampilkanCart", func(t *testing.T) {
		filter := entitas.Filter{
			UserId:     1,
			NamaProduk: "product 1",
			Kuantitas:  1,
		}

		cart, err := repo.TampilkanCart(filter)
		if err != nil {
			t.Errorf("error while filtering cart: %v", err)
		}

		if len(cart.Products) != 1 {
			t.Errorf("error while filtering cart, expected: %d, got: %d", 1, len(cart.Products))
		}

	})
}

type mockDBCart struct {
	dataCart []entitas.Cart
}

func (m *mockDBCart) Load(key string) ([]byte, error) {
	bytes, err := json.Marshal(&m.dataCart)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (m *mockDBCart) Save(key string, data []byte) error {
	return nil
}

type mockRepoProduct struct {
	dataProduct []entitas.Product
}

func (p *mockRepoProduct) LihatProduk() ([]entitas.Product, error) {
	return p.dataProduct, nil
}

func (p *mockRepoProduct) CariProduk(kode string) (*entitas.Product, error) {
	for _, val := range p.dataProduct {
		if val.KodeProduk == kode {
			return &val, nil
		}
	}

	return nil, errors.New("product not found")
}

func (p *mockRepoProduct) CariIndex(kodeProduk string) bool {
	for _, val := range p.dataProduct {
		if val.KodeProduk == kodeProduk {
			return true
		}
	}

	return false
}
