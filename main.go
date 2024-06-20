package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Dvd struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
}

var dvds = []Dvd{
	{ID: "1", Title: "Family Guy", Price: 13.99},
	{ID: "2", Title: "The Sopranos", Price: 22.99},
	{ID: "3", Title: "The Wire",Price: 14.99},
}

func getSales(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, dvds)
}

func main() {
	router := gin.Default()
	router.GET("/dvds", getSales)
	router.Run("localhost:8080")
}


