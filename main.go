package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// No Route
	router.NoRoute(noRoute)

	// Routes
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// To keep things simple for the tutorial, you’ll store data in memory. A more typical API would interact with a database.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// only for this tutorial the data will be stored in memory and will be deleted with each new run.
// for a real application use a database for example.
var albums = []album{
	{ID: "1", Title: "Blue train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func noRoute(c *gin.Context) {
	// render a page not found here.

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "page not found"})
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	// in a real application you should get data from a database by example.
	c.IndentedJSON(http.StatusOK, albums)

	// use this to minimize the json. the Idented is for debugging.
	// c.JSON(http.StatusOK, albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// a real-world service would likely use a database query to perform this lookup.
	// loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// call bindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// in a real application you must put the data in a database by example.
	// add the new album to the slice.
	albums = append(albums, newAlbum)

	// show the new album added.
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
