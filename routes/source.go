package routes

import (
	"assignment/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func findSourceByID(sources *[]models.Source, id string) (*models.Source, int) {
	for i, s := range *sources {
		if s.ID == id {
			return &(*sources)[i], i
		}
	}
	return nil, -1
}

func RegisterSourceRoutes(r *gin.Engine, sources *[]models.Source) {
	r.GET("/sources", getSources(sources))
	r.GET("/sources/:id", getSourceByID(sources))
	r.POST("/sources", createSource(sources))
	r.PUT("/sources/:id", updateSource(sources))
	r.DELETE("/sources/:id", deleteSource(sources))
}

func getSources(sources *[]models.Source) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, models.StandardResponse{Message: "Sources fetched successfully", Data: *sources})
	}
}

func getSourceByID(sources *[]models.Source) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		source, _ := findSourceByID(sources, id)

		if source == nil {
			c.JSON(http.StatusNotFound, models.StandardResponse{Error: "Source not found"})
			return
		}
		c.JSON(http.StatusOK, models.StandardResponse{Message: "Source fetched successfully", Data: *source})
	}
}

func createSource(sources *[]models.Source) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newSource models.Source
		if err := c.ShouldBindJSON(&newSource); err != nil {
			c.JSON(http.StatusBadRequest, models.StandardResponse{Error: err.Error()})
			return
		}
		newSource.ID = uuid.New().String()
		*sources = append(*sources, newSource)
		c.JSON(http.StatusCreated, models.StandardResponse{Message: "Source created successfully", Data: newSource})
	}
}

func updateSource(sources *[]models.Source) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		_, index := findSourceByID(sources, id)

		if index == -1 {
			c.JSON(http.StatusNotFound, models.StandardResponse{Error: "Source not found"})
			return
		}

		var updatedSourceData models.Source
		if err := c.ShouldBindJSON(&updatedSourceData); err != nil {
			c.JSON(http.StatusBadRequest, models.StandardResponse{Error: err.Error()})
			return
		}

		(*sources)[index].Name = updatedSourceData.Name

		c.JSON(http.StatusOK, models.StandardResponse{Message: "Source updated successfully", Data: (*sources)[index]})
	}
}

func deleteSource(sources *[]models.Source) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		_, index := findSourceByID(sources, id)

		if index == -1 {
			c.JSON(http.StatusNotFound, models.StandardResponse{Error: "Source not found"})
			return
		}

		*sources = append((*sources)[:index], (*sources)[index+1:]...)
		c.JSON(http.StatusOK, models.StandardResponse{Message: "Source deleted successfully"})
	}
}