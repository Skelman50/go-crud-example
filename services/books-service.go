package services

import (
	db "books-rest-api/config"
	m "books-rest-api/models"
	"fmt"
	"strings"
)

func FindBookOnce(id string) []m.BookModel {
	rows, err := db.Db.Query("SELECT * FROM myTestDB.books where id=?", id)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	books := []m.BookModel{}
	for rows.Next() {
		b := m.BookModel{}
		err := rows.Scan(&b.ID, &b.Title, &b.Author)
		if err != nil {
			fmt.Println(err)
			continue
		}
		books = append(books, b)
	}
	return books
}

func FindAllBooks() ([]m.BookModel, bool) {
	rows, err := db.Db.Query("SELECT * FROM myTestDB.books")
	if err != nil {
		return nil, true
	}
	defer rows.Close()
	books := []m.BookModel{}
	for rows.Next() {
		b := m.BookModel{}
		err := rows.Scan(&b.ID, &b.Title, &b.Author)
		if err != nil {
			fmt.Println(err)
			continue
		}
		books = append(books, b)
	}
	return books, false
}

func UpdateBookService(id string, book m.BookModel) {
	if book.Author != nil && book.Title != nil {
		_, err := db.Db.Exec("update myTestDB.books set title = ?, author = ? where id = ?", book.Title, book.Author, id)
		if err != nil {
			panic(err)
		}
	}

	if book.Author != nil && book.Title == nil {
		_, err := db.Db.Exec("update myTestDB.books set author = ? where id = ?", book.Author, id)
		if err != nil {
			panic(err)
		}
	}

	if book.Author == nil && book.Title != nil {
		_, err := db.Db.Exec("update myTestDB.books set title = ? where id = ?", book.Title, id)
		if err != nil {
			panic(err)
		}
	}
}

func InsertBookService(book m.PostBook) {
	_, err := db.Db.Exec("insert into myTestDB.books (title, author) values (?, ?)",
		book.Title, book.Author)
	if err != nil {
		panic(err)
	}
}

func DeleteBookService(id string) (int64, strings.Builder) {
	result, err := db.Db.Exec("delete FROM myTestDB.books where id = ?", id)
	if err != nil {
		panic(err)
	}

	rows, _ := result.RowsAffected()
	var str strings.Builder
	buffer := []string{"book with id ", id, " not found"}
	for _, s := range buffer {
		str.WriteString(s)
	}

	return rows, str
}
