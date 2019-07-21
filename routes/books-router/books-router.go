package router

import (
	ctrls "books-rest-api/controllers/books-controllers"

	"github.com/gin-gonic/gin"
)

func BooksRouter(v1 *gin.RouterGroup) {
	{
		v1.GET("/books", ctrls.QueryAllBooks)
		v1.GET("/book/:id", ctrls.FindOneBook)
		v1.PUT("/book/:id", ctrls.UpdateBook)
		v1.DELETE("/book/:id", ctrls.DeleteBook)
		v1.POST("/book", ctrls.InsertNewBook)
	}
}
