package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	r := gin.Default()
	r.Use(myMiddleware())

	api := r.Group("/api")
	{

		v1 := api.Group("/v1")
		{
			v1.Use(myGroupV1Middleware())
			v1.GET("/ping", getData)
			v1.GET("/detail/:id", getDetailData)
		}
		v2 := api.Group("/v2")
		{
			v2.POST("/Postping", postData)
			v2.POST("/PostJSON", postDataJSON)
		}

	}

	r.Run(":8081")
}

func getDetailData(context *gin.Context) {
	//ID := context.Param("id")
	//context.JSON(http.StatusOK, gin.H{
	//	"id": ID,
	//})

	id := context.Param("id")
	for _, a := range albums {
		if a.ID == id {
			context.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func getData(context *gin.Context) {

	//name := context.DefaultQuery("name", "guest")
	//context.JSON(http.StatusOK, gin.H{
	//	"message": "Hello " + name + " from Y",
	//})
	context.IndentedJSON(http.StatusOK, albums)
}

func postData(context *gin.Context) {
	//address := context.DefaultPostForm("addr", "VietNam")
	//context.JSON(http.StatusOK, gin.H{
	//	"message": "Hello " + address + " from POST ",
	//})

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	var newAlbum album
	if err := context.BindJSON(&newAlbum); err != nil {
		return
	}
	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, newAlbum)

}

func postDataJSON(context *gin.Context) {
	var newAlbum album
	context.BindJSON(&newAlbum)
	albums = append(albums, newAlbum)
	context.JSON(200, newAlbum)
}

//func myMiddleware(context *gin.Context) {
//	log.Println("I'm global middleware")
//	context.Next()
//}

func myMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Println("I'm global middleware")
		context.Next()
	}
}

func myGroupV1Middleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Println("I'm groupV1 middleware")
		context.Next()
	}
}
