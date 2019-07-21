package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type BookModel struct {
	ID     int64   `json:"id"`
	Title  *string `*json:"title"`
	Author *string `*json:"author"`
}

type PostBook struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (p PostBook) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Author, validation.Required),
		validation.Field(&p.Title, validation.Required),
	)
}
