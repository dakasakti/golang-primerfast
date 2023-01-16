package repositories

import (
	"encoding/json"
	"testing"

	"github.com/dakasakti/golang/primerfast/entitas"
)

func TestProduk(t *testing.T) {
	mockDB := &mockDB{
		data: []entitas.Product{
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
		},
	}

	repo := NewRepoProduct(mockDB)

	t.Run("lihatProduk", func(t *testing.T) {
		products, err := repo.LihatProduk()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(products) != 2 {
			t.Errorf("Expected 2 products, got %v", len(products))
		}

		if products[0].KodeProduk != "DM-001" {
			t.Errorf("Expected product with code DM-001, got %v", products[0].KodeProduk)
		}

		if products[1].KodeProduk != "DM-002" {
			t.Errorf("Expected product with code DM-002, got %v", products[1].KodeProduk)
		}
	})

	t.Run("cariProduk", func(t *testing.T) {
		product, err := repo.CariProduk("DM-001")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if product.KodeProduk != "DM-001" {
			t.Errorf("Expected product with code DM-001, got %v", product.KodeProduk)
		}

		product, err = repo.CariProduk("DM-003")
		if err == nil {
			t.Error("Expected an error, got nil")
		}
		if product != nil {
			t.Errorf("Expected nil product, got %v", product)
		}
	})
}

type mockDB struct {
	data []entitas.Product
}

func (m *mockDB) Load(key string) ([]byte, error) {
	bytes, err := json.Marshal(&m.data)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (m *mockDB) Save(key string, data []byte) error {
	return nil
}
