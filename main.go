package main

import (
	"beauty-shop/api"
	"beauty-shop/models"
	repo "beauty-shop/repository"
	serv "beauty-shop/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func initDB() *gorm.DB {
	dsn := "host=localhost " + "user=" + os.Getenv("APP_DB_USERNAME") + " password=" + os.Getenv("APP_DB_PASSWORD") +
		" dbname=" + os.Getenv("APP_DB_NAME") + " sslmode=disable port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	_ = db.AutoMigrate(&models.Book{})

	return db
}

func InitProductApi(db *gorm.DB) api.BookApi {
	repository := repo.ProvideBookRepository(db)
	service := serv.ProvideBookService(repository)
	api2 := api.ProviderBookApi(service)
	return api2
}

func main() {
	db := initDB()

	bookAPI := InitProductApi(db)

	r := gin.Default()

	r.GET("/books", bookAPI.FindAll)
	r.GET("/book/:id", bookAPI.FindByID)
	r.POST("/book", bookAPI.Create)
	r.PUT("/book/:id", bookAPI.Update)
	r.DELETE("/book/:id", bookAPI.Delete)

	err := r.Run(":8010")
	if err != nil {
		panic(err)
	}
}
