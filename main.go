package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	TITLE  string  `json:"title"`
	ARTIST string  `json:"artist"`
	PRICE  float64 `json:"price"`
}

var albums = []album{
	{
		ID: "1", TITLE: "BLUE Train", ARTIST: "JOHN Coltrane", PRICE: 56.99,
	},
	{
		ID: "2", TITLE: "JERRU", ARTIST: "Gerry Mulligan", PRICE: 17.99,
	},
}

func get_albums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", get_albums)

	router.Run("localhost:8080")
}
