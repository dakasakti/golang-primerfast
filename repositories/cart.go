package repositories

import (
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/dakasakti/golang/primerfast/db"
	"github.com/dakasakti/golang/primerfast/entitas"
)

type repoCart struct {
	db.DB
	RepoProduct
}

func NewRepoCart(db db.DB, product RepoProduct) *repoCart {
	return &repoCart{db, product}
}

func (p *repoCart) lihatCart() ([]entitas.Cart, error) {
	bytes, err := p.DB.Load("carts")
	if err != nil {
		return nil, err
	}

	var results []entitas.Cart
	err = json.Unmarshal(bytes, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (p *repoCart) cariCart(userId int) (*entitas.Cart, error) {
	carts, err := p.lihatCart()
	if err != nil {
		return nil, err
	}

	for i, val := range carts {
		if val.UserID == userId {
			return &carts[i], nil
		}
	}

	return nil, errors.New("cart not found")
}

func (p *repoCart) TampilkanCart(filter entitas.Filter) (*entitas.Cart, error) {
	cart, err := p.cariCart(filter.UserId)
	if err != nil {
		return nil, err
	}

	if filter.NamaProduk == "" && filter.Kuantitas == 0 {
		return cart, nil
	}

	var filteredProducts []entitas.Product
	for _, val := range cart.Products {
		if (strings.Contains(strings.ToLower(val.NamaProduk), strings.ToLower(filter.NamaProduk))) && (val.Kuantitas == filter.Kuantitas) {
			log.Printf("%s - %s - (%d)\n", val.KodeProduk, val.NamaProduk, val.Kuantitas)
			filteredProducts = append(filteredProducts, val)
		} else if (filter.NamaProduk != "") && strings.Contains(strings.ToLower(val.NamaProduk), strings.ToLower(filter.NamaProduk)) {
			log.Printf("%s - %s - (%d)\n", val.KodeProduk, val.NamaProduk, val.Kuantitas)
			filteredProducts = append(filteredProducts, val)
		} else if val.Kuantitas == filter.Kuantitas {
			log.Printf("%s - %s - (%d)\n", val.KodeProduk, val.NamaProduk, val.Kuantitas)
			filteredProducts = append(filteredProducts, val)
		}
	}

	cart.Products = filteredProducts
	return cart, nil
}

func (p *repoCart) TambahCart(Userid int, data entitas.CartRequest) error {
	res := p.CariIndex(data.KodeProduk)
	if !res {
		return errors.New("the product doesn't exist")
	}

	carts, err := p.lihatCart()
	if err != nil {
		return err
	}

	var cart *entitas.Cart
	var index int
	var exist bool

	for i, val := range carts {
		if val.UserID == Userid {
			cart = &val
			index = i
			exist = true
			break
		}
	}

	if !exist {
		cart = &entitas.Cart{
			UserID: Userid,
		}

		carts = append(carts, *cart)
		index = len(carts) - 1
	}

	product, err := p.CariProduk(data.KodeProduk)
	if err != nil {
		return err
	}

	product.Kuantitas = data.Kuantitas

	exist = false
	for i, val := range cart.Products {
		if val.KodeProduk == data.KodeProduk {
			cart.Products[i].Kuantitas += data.Kuantitas
			exist = true
			break
		}
	}

	if !exist {
		cart.Products = append(cart.Products, *product)
	}

	carts[index] = *cart

	bytes, err := json.Marshal(&carts)
	if err != nil {
		return err
	}

	err = p.DB.Save("carts", bytes)
	if err != nil {
		return err
	}

	return nil
}

func (p *repoCart) HapusProdukCart(Userid int, kodeProduk string) error {
	carts, err := p.lihatCart()
	if err != nil {
		return err
	}

	var cart *entitas.Cart
	var index int

	for i, val := range carts {
		if val.UserID == Userid {
			cart = &val
			index = i
			break
		}
	}

	if cart == nil {
		return errors.New("cart with this user id not found")
	}

	for i, val := range cart.Products {
		if val.KodeProduk == kodeProduk {
			cart.Products = append(cart.Products[:i], cart.Products[i+1:]...)
			break
		}
	}

	carts[index] = *cart

	bytes, err := json.Marshal(&carts)
	if err != nil {
		return err
	}

	err = p.DB.Save("carts", bytes)
	if err != nil {
		return err
	}

	return nil
}
