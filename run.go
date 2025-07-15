package main

import (
	"assignment/models"
	"assignment/routes"

	"github.com/gin-gonic/gin"
)

var (
	products     []models.Product
	sources      []models.Source
	transactions []models.Transaction
)

func main() {
	sources = []models.Source{
		{ID: "s1", Name: "IMDHED"},
		{ID: "s2", Name: "IMPHEN"},
	}
	products = []models.Product{
		{ID: "p1", Name: "Wallpaper", Description: "Wallpaper HD++", Price: 150000, Stock: 20, SourceID: "s1"},
		{ID: "p2", Name: "Game Mobile", Description: "Game Mobile terbaru buatan IMPHEN", Price: 25000, Stock: 100, SourceID: "s2"},
	}

	r := gin.Default()

	routes.RegisterProductRoutes(r, &products, &sources)
	routes.RegisterSourceRoutes(r, &sources)
	routes.RegisterTransactionRoutes(r, &transactions, &products)

	r.Run(":8080")
}