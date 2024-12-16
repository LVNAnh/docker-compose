package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items = []Item{
	{ID: "1", Name: "Item 1"},
	{ID: "2", Name: "Item 2"},
}

func getItems(c *gin.Context) {
	c.JSON(http.StatusOK, items)
}

func createItem(c *gin.Context) {
	var newItem Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	items = append(items, newItem)
	c.JSON(http.StatusOK, newItem)
}

func main() {
	router := gin.Default()

	router.GET("/items", getItems)
	router.POST("/create-items", createItem)

	router.Run(":8080")
}
