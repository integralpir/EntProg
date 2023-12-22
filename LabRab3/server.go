package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type productType struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ImageURL    string `json:"img_url"`
}

var products = []productType{
	{
		Name:        "Товар 1",
		Description: "Описание товара 1",
		Price:       100,
		ImageURL:    "/assets/img/article1.png",
	},
	{
		Name:        "Товар 2",
		Description: "Описание товара 2",
		Price:       200,
		ImageURL:    "/assets/img/article2.png",
	},
	{
		Name:        "Товар 3",
		Description: "Описание товара 3",
		Price:       300,
		ImageURL:    "/assets/img/article3.png",
	},
	{
		Name:        "Товар 4",
		Description: "Описание товара 4",
		Price:       400,
		ImageURL:    "/assets/img/article1.png",
	},
	{
		Name:        "Товар 5",
		Description: "Описание товара 5",
		Price:       500,
		ImageURL:    "/assets/img/article2.png",
	},
	{
		Name:        "Товар 6",
		Description: "Описание товара 6",
		Price:       600,
		ImageURL:    "/assets/img/article3.png",
	},
	{
		Name:        "Товар 7",
		Description: "Описание товара 7",
		Price:       700,
		ImageURL:    "/assets/img/article2.png",
	},
	{
		Name:        "Товар 8",
		Description: "Описание товара 8",
		Price:       800,
		ImageURL:    "/assets/img/article3.png",
	},
}

const productPerPage = 3

func main() {
	e := echo.New()

	e.Static("/public", "public")
	e.Static("/assets", "public/assets")

	e.POST("/getProducts", func(c echo.Context) error {
		type requestData struct {
			Page int `json:"page"`
		}
		var request requestData
		err := c.Bind(&request)

		if err != nil {
			return err
		}

		startIndex := (request.Page - 1) * productPerPage
		endIndex := startIndex + productPerPage

		if endIndex > len(products) {
			endIndex = len(products)
		}

		return c.JSON(http.StatusOK, products[startIndex:endIndex])
	})

	e.POST("/getProductCount", func(c echo.Context) error {
		type responseResult = struct {
			Count           int `json:"count"`
			ProductsPerPage int `json:"products_per_page"`
		}

		result := responseResult{
			Count:           len(products),
			ProductsPerPage: productPerPage,
		}

		return c.JSON(http.StatusOK, result)
	})

	e.GET("*", func(c echo.Context) error {
		return c.File("index.html")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
