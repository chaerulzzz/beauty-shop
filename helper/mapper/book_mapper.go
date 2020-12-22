package mapper

import (
	"beauty-shop/dto"
	"beauty-shop/models"
)

func ToBook(dto dto.BookDTO) models.Book {
	return models.Book{Title: dto.Title, Author: dto.Author}
}

func ToBookDTO(book models.Book) dto.BookDTO {
	return dto.BookDTO{ID: book.ID, Title: book.Title, Author: book.Author}
}

func ToBookDTOs(books []models.Book) []dto.BookDTO {
	booksDTOs := make([]dto.BookDTO, len(books))

	for i, item := range books {
		booksDTOs[i] = ToBookDTO(item)
	}

	return booksDTOs
}
