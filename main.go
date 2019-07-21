package main

import (
	db "books-rest-api/config"
	"books-rest-api/helpers"
	booksRouter "books-rest-api/routes/books-router"

	"github.com/gin-gonic/gin"
)

func main() {
	helpers.MaxOpenFiles()
	db.InitDB()
	defer db.Db.Close()
	r := gin.New()
	r.Use(gin.Recovery())
	v1 := r.Group("/api/v1")
	booksRouter.BooksRouter(v1)
	r.Run()
}
