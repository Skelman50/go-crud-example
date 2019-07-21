package controllers

import (
	m "books-rest-api/models"
	s "books-rest-api/services"

	"github.com/gin-gonic/gin"
)

func QueryAllBooks(c *gin.Context) {
	books, err := s.FindAllBooks()
	if err == true {
		c.JSON(200, gin.H{"error": "error"})
		return
	}
	c.JSON(200, gin.H{"books": books})
}

func FindOneBook(c *gin.Context) {
	id := c.Param("id")
	books := s.FindBookOnce(id)
	if len(books) <= 0 {
		c.JSON(404, gin.H{"error": "book not found"})
		return
	}
	c.JSON(200, gin.H{"books": books})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	book := m.BookModel{}
	c.BindJSON(&book)
	if c.Writer.Status() == 400 {
		c.JSON(400, gin.H{"error": "bad request"})
		return
	}

	if book.Title == nil && book.Author == nil {
		c.JSON(400, gin.H{"error": "bad request"})
		return
	}

	s.UpdateBookService(id, book)

	c.JSON(200, gin.H{
		"message": "updated",
	})
}

func InsertNewBook(c *gin.Context) {
	book := m.PostBook{}
	c.BindJSON(&book)
	notValid := book.Validate()
	if notValid != nil {
		c.JSON(c.Writer.Status(), gin.H{"error": notValid.Error()})
		return
	}
	s.InsertBookService(book)
	c.JSON(201, gin.H{"status": "created"})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	rows, str := s.DeleteBookService(id)
	if rows <= 0 {
		c.JSON(404, gin.H{"books": str.String()})
		return
	}
	c.JSON(200, gin.H{"books": "book deleted"})
}
