package requests

import "miniproject/business/books"

type SearchBook struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

func ToDomain(search SearchBook) books.Domain {
	return books.Domain{
		Title: 		search.Title,
		Author: 	search.Author,
		Publisher: 	search.Publisher,
	}
}