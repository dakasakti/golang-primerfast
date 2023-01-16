package repositories

import (
	"encoding/json"
	"errors"

	"github.com/dakasakti/golang/primerfast/db"
	"github.com/dakasakti/golang/primerfast/entitas"
)

type RepoProduct interface {
	LihatProduk() ([]entitas.Product, error)
	CariProduk(kode string) (*entitas.Product, error)
	CariIndex(kodeProduk string) bool
}

type repoProduct struct {
	db.DB
}

func NewRepoProduct(db db.DB) *repoProduct {
	return &repoProduct{db}
}

func (p *repoProduct) LihatProduk() ([]entitas.Product, error) {
	bytes, err := p.DB.Load("products")
	if err != nil {
		return nil, err
	}

	var results []entitas.Product

	err = json.Unmarshal(bytes, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (p *repoProduct) CariProduk(kode string) (*entitas.Product, error) {
	products, err := p.LihatProduk()
	if err != nil {
		return nil, err
	}

	for i, val := range products {
		if val.KodeProduk == kode {
			return &products[i], nil
		}
	}

	return nil, errors.New("the product not found")
}

func (p *repoProduct) TambahProduk(data entitas.Product) error {
	res := p.CariIndex(data.KodeProduk)
	if res {
		return errors.New("the product already exist")
	}

	old, err := p.LihatProduk()
	if err != nil {
		return err
	}

	old = append(old, data)

	bytes, err := json.Marshal(&old)
	if err != nil {
		return err
	}

	err = p.DB.Save("products", bytes)
	if err != nil {
		return err
	}

	err = p.tambahIndex(data.KodeProduk)
	if err != nil {
		return err
	}

	return nil
}

func (p *repoProduct) tambahIndex(kodeProduk string) error {
	old, err := p.lihatIndex()
	if err != nil {
		return err
	}

	old[kodeProduk] = true

	bytes, err := json.Marshal(&old)
	if err != nil {
		return err
	}

	err = p.DB.Save("products_index", bytes)
	if err != nil {
		return err
	}

	return nil
}

func (p *repoProduct) lihatIndex() (map[string]bool, error) {
	bytes, err := p.DB.Load("products_index")
	if err != nil {
		return nil, err
	}

	var results map[string]bool
	err = json.Unmarshal(bytes, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (p *repoProduct) CariIndex(kodeProduk string) bool {
	results, err := p.lihatIndex()
	if err != nil {
		return false
	}

	return results[kodeProduk]
}
