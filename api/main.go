package main

import(
	"net/http"
	// "errors"
	"github.com/gin-gonic/gin"
)

type book struct{
	ID  	 string `json:"id"`
	Author   string `json:"author"`
	Title    string `json:"title"`
	Quantity string `json:"quantity"`
}

var books = []book{
	{ID: "1", Author: "Chris Guillebeau", Title:"The 100$ startup", Quantity:"8"},
	{ID: "2", Author: "Atul Subedi", Title:"The nonsense shit", Quantity: "9"},
	{ID: "3", Author: "Atulya Subedi", Title:"The art of moving", Quantity: "4" },
}

func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, books)
}

func createBooks (c *gin.Context){
	var newBooks book

	if err := c.BindJSON(&newBooks); err != nil{
		return
	}

	books = append(books, newBooks)
	c.IndentedJSON(http.StatusCreated, newBooks)
}

func main(){
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBooks)
	router.Run("localhost:8080")
}
