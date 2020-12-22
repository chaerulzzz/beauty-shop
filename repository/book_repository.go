package repository

import (
	"beauty-shop/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func ProvideBookRepository(DB *gorm.DB) BookRepository {
	return BookRepository{DB: DB}
}

func (p *BookRepository) FindAll() []models.Book {
	var books []models.Book
	p.DB.Find(&books)

	return books
}

func (p *BookRepository) FindByID(id uint) models.Book {
	var book models.Book
	p.DB.Find(&book, id)

	return book
}

func (p *BookRepository) Save(book models.Book) models.Book {
	p.DB.Save(&book)

	return book
}

func (p *BookRepository) Delete(book models.Book) {
	p.DB.Delete(&book)
}
