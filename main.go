package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/dakasakti/golang/primerfast/db"
	"github.com/dakasakti/golang/primerfast/entitas"
	"github.com/dakasakti/golang/primerfast/feature"
	"github.com/dakasakti/golang/primerfast/helper"
	"github.com/dakasakti/golang/primerfast/repositories"
	"github.com/labstack/echo/v4"
)

func main() {
	db := &db.JsonDB{}
	product := repositories.NewRepoProduct(db)
	cart := repositories.NewRepoCart(db, product)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "System API design by Mahmuda Karima (Daka)")
	})

	e.POST("/questions/first", func(c echo.Context) error {
		var req entitas.CountMoneyRequest

		err := c.Bind(&req)
		if err != nil {
			log.Println("[POST] /questions/first", err.Error())
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "failed binding data",
			})
		}

		if req.Nominal == "" {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "failed requestBody",
			})
		}

		newReq := strings.Replace(req.Nominal, ".", "", -1)
		data, err := strconv.ParseInt(newReq, 10, 64)
		if err != nil {
			log.Println("[POST] /questions/first", err.Error())
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "failed convert data",
			})
		}

		res := feature.CountMoney(data)
		return c.JSON(http.StatusCreated, res)
	})

	e.POST("/questions/second", func(c echo.Context) error {
		var req entitas.BooleanDataRequest

		err := c.Bind(&req)
		if err != nil {
			log.Println("[POST] /questions/second", err.Error())
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "failed binding data",
			})
		}

		if req.Except == "" || req.Input == "" {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "failed requestBody",
			})
		}

		res := feature.BooleanData(req.Except, req.Input)
		return c.JSON(http.StatusCreated, echo.Map{
			"data": res,
		})
	})

	e.POST("/products", func(c echo.Context) error {
		var req entitas.Product

		err := c.Bind(&req)
		if err != nil {
			log.Println("[POST] /products", err.Error())
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "failed binding data",
			})
		}

		if req.KodeProduk == "" || req.NamaProduk == "" || req.Kuantitas == 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "failed requestBody",
			})
		}

		data := entitas.Product{
			KodeProduk: req.KodeProduk,
			NamaProduk: req.NamaProduk,
			Kuantitas:  req.Kuantitas,
		}

		err = product.TambahProduk(data)
		if err != nil {
			log.Println("[POST] /products", err.Error())
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, echo.Map{
			"message": "success created data",
		})
	})

	e.GET("/products", func(c echo.Context) error {
		res, err := product.LihatProduk()
		if err != nil {
			log.Println("[GET] /products", err.Error())
			return c.JSON(500, echo.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, res)
	})

	e.POST("/carts", func(c echo.Context) error {
		userId := c.Request().Header.Get("Authorization")

		user_id, err := helper.ConvertInt(userId)
		if err != nil {
			log.Println("[GET] /carts", err.Error())
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "failed access data user",
			})
		}

		var req entitas.CartRequest

		err = c.Bind(&req)
		if err != nil {
			log.Println("[POST] /carts", err.Error())
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "failed binding data",
			})
		}

		if req.KodeProduk == "" || req.Kuantitas == 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "failed requestBody",
			})
		}

		err = cart.TambahCart(user_id, req)
		if err != nil {
			log.Println("[POST] /carts", err.Error())
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, echo.Map{
			"message": "success added product to carts",
		})
	})

	e.GET("/carts", func(c echo.Context) error {
		userId := c.Request().Header.Get("Authorization")

		user_id, err := helper.ConvertInt(userId)
		if err != nil {
			log.Println("[GET] /carts", err.Error())
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "failed access data user",
			})
		}

		namaProduk := c.QueryParam("namaProduk")

		filter := entitas.Filter{
			UserId:     user_id,
			NamaProduk: namaProduk,
		}

		if queryKuantitas := c.QueryParam("kuantitas"); queryKuantitas != "" {
			kuantitas, err := helper.ConvertInt(queryKuantitas)
			if err != nil {
				log.Println("[GET] /carts", err.Error())
				return c.JSON(http.StatusBadRequest, echo.Map{
					"message": "failed convert kuantitas",
				})
			}

			filter.Kuantitas = kuantitas
		}

		res, err := cart.TampilkanCart(filter)
		if err != nil {
			log.Println("[GET] /carts", err.Error())
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, res)
	})

	e.DELETE("/carts/:id", func(c echo.Context) error {
		user_id, err := helper.ConvertInt(c.Request().Header.Get("Authorization"))
		if err != nil {
			log.Println("[DELETE] /carts", err.Error())
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "failed access data user",
			})
		}

		kodeProduk := c.Param("id")
		err = cart.HapusProdukCart(user_id, kodeProduk)
		if err != nil {
			log.Println("[DELETE] /carts", err.Error())
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "success deleted product from cart",
		})
	})

	e.Logger.Fatal(e.Start(":3000"))
}
