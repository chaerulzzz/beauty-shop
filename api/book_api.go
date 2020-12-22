package api

import (
	"beauty-shop/dto"
	"beauty-shop/helper/mapper"
	"beauty-shop/models"
	"beauty-shop/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BookApi struct {
	BookService service.BookService
}

func ProviderBookApi(p service.BookService) BookApi {
	return BookApi{BookService: p}
}

func (p *BookApi) FindAll(c *gin.Context) {
	books := p.BookService.FindAll()

	if len(books) <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"books":   mapper.ToBookDTOs(books),
			"message": "List is empty!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"books": mapper.ToBookDTOs(books)})
}

func (p *BookApi) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Book ID"})
		return
	}

	book := p.BookService.FindByID(uint(id))
	if book == (models.Book{}) {
		c.JSON(http.StatusOK, gin.H{
			"book":    "{}",
			"message": "Book not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"book": mapper.ToBookDTO(book)})
}

func (p *BookApi) Create(c *gin.Context) {
	var bookDTO dto.BookDTO
	err := c.BindJSON(&bookDTO)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	createBook := p.BookService.Save(mapper.ToBook(bookDTO))

	c.JSON(http.StatusCreated, gin.H{"book": mapper.ToBookDTO(createBook)})
}

func (p *BookApi) Update(c *gin.Context) {
	var bookDTO dto.BookDTO
	err := c.BindJSON(&bookDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Book ID"})
		return
	}

	book := p.BookService.FindByID(uint(id))
	if book == (models.Book{}) {
		c.JSON(http.StatusOK, gin.H{"message": "Book not found!"})
		return
	}

	book.Author = bookDTO.Author
	book.Title = bookDTO.Title
	p.BookService.Save(book)

	c.JSON(http.StatusNoContent, gin.H{"message": ""})
}

func (p *BookApi) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Book ID"})
		return
	}

	book := p.BookService.FindByID(uint(id))
	if book == (models.Book{}) {
		c.JSON(http.StatusOK, gin.H{
			"book":    "{}",
			"message": "Book not found",
		})
		return
	}

	p.BookService.Delete(book)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
