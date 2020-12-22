package service

import (
	"beauty-shop/models"
	"beauty-shop/repository"
)

type BookService struct {
	BookRepository repository.BookRepository
}

func ProvideBookService(p repository.BookRepository) BookService {
	return BookService{BookRepository: p}
}

func (p *BookService) FindAll() []models.Book {
	return p.BookRepository.FindAll()
}

func (p *BookService) FindByID(id uint) models.Book {
	return p.BookRepository.FindByID(id)
}

func (p *BookService) Save(book models.Book) models.Book {
	p.BookRepository.Save(book)

	return book
}

func (p *BookService) Delete(book models.Book) {
	p.BookRepository.Delete(book)
}
