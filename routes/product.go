package routes

import (
	"assignment/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func findProductByID(products *[]models.Product, id string) (*models.Product, int) {
	for i, p := range *products {
		if p.ID == id {
			return &(*products)[i], i
		}
	}
	return nil, -1
}

func RegisterProductRoutes(r *gin.Engine, products *[]models.Product, sources *[]models.Source) {
	r.GET("/products", getProducts(products))
	r.GET("/products/:id", getProductByID(products))
	r.POST("/products", createProduct(products, sources))
	r.PUT("/products/:id", updateProduct(products, sources))
	r.DELETE("/products/:id", deleteProduct(products))
}

func getProducts(products *[]models.Product) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, models.StandardResponse{Message: "Products fetched successfully", Data: *products})
	}
}

func getProductByID(products *[]models.Product) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		product, _ := findProductByID(products, id)

		if product == nil {
			c.JSON(http.StatusNotFound, models.StandardResponse{Error: "Product not found"})
			return
		}
		c.JSON(http.StatusOK, models.StandardResponse{Message: "Product fetched successfully", Data: *product})
	}
}

func createProduct(products *[]models.Product, sources *[]models.Source) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newProduct models.Product
		if err := c.ShouldBindJSON(&newProduct); err != nil {
			c.JSON(http.StatusBadRequest, models.StandardResponse{Error: err.Error()})
			return
		}

		sourceExists := false
		for _, s := range *sources {
			if s.ID == newProduct.SourceID {
				sourceExists = true
				break
			}
		}
		if !sourceExists {
			c.JSON(http.StatusBadRequest, models.StandardResponse{Error: "Invalid source_id"})
			return
		}

		newProduct.ID = uuid.New().String()
		*products = append(*products, newProduct)
		c.JSON(http.StatusCreated, models.StandardResponse{Message: "Product created successfully", Data: newProduct})
	}
}

func updateProduct(products *[]models.Product, sources *[]models.Source) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		_, index := findProductByID(products, id)

		if index == -1 {
			c.JSON(http.StatusNotFound, models.StandardResponse{Error: "Product not found"})
			return
		}

		var updatedProductData models.Product
		if err := c.ShouldBindJSON(&updatedProductData); err != nil {
			c.JSON(http.StatusBadRequest, models.StandardResponse{Error: err.Error()})
			return
		}

		sourceExists := false
		for _, s := range *sources {
			if s.ID == updatedProductData.SourceID {
				sourceExists = true
				break
			}
		}
		if !sourceExists {
			c.JSON(http.StatusBadRequest, models.StandardResponse{Error: "Invalid source_id"})
			return
		}

		(*products)[index].Name = updatedProductData.Name
		(*products)[index].Description = updatedProductData.Description
		(*products)[index].Price = updatedProductData.Price
		(*products)[index].Stock = updatedProductData.Stock
		(*products)[index].SourceID = updatedProductData.SourceID

		c.JSON(http.StatusOK, models.StandardResponse{Message: "Product updated successfully", Data: (*products)[index]})
	}
}

func deleteProduct(products *[]models.Product) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		_, index := findProductByID(products, id)

		if index == -1 {
			c.JSON(http.StatusNotFound, models.StandardResponse{Error: "Product not found"})
			return
		}

		*products = append((*products)[:index], (*products)[index+1:]...)
		c.JSON(http.StatusOK, models.StandardResponse{Message: "Product deleted successfully"})
	}
}