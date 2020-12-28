package main

import (
	"beauty-shop/auth"
	"beauty-shop/delivery"
	repo "beauty-shop/repository"
	serv "beauty-shop/usecase"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func InitProductApi(base repo.BaseRepository) delivery.BookHandler {
	repository := repo.ProvideBookRepository(base)
	service := serv.ProvideBookUseCase(repository)
	api := delivery.ProviderBookHandler(service)
	return api
}

func InitUserApi(base repo.BaseRepository) delivery.UserHandler {
	repository := repo.ProvideUserRepository(base)
	service := serv.ProvideUserUseCase(repository)
	api := delivery.ProviderUserHandler(service)

	return api
}

func initDB() repo.BaseRepository {
	dbDriver := os.Getenv("APP_DB_NAME")
	username := os.Getenv("APP_DB_USERNAME")
	password := os.Getenv("APP_DB_PASSWORD")
	host := os.Getenv("APP_DB_HOST")
	database := os.Getenv("APP_DB_NAME")
	dbPort := os.Getenv("APP_DB_PORT")

	repos, err := repo.Model.Initialize(dbDriver, username, password, dbPort, host, database)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	return repos
}

func initRouter(base repo.BaseRepository) *gin.Engine {
	r := gin.Default()
	main := r.Group("beauty-shop")
	setBookRouter(main, base)
	setLoginRegister(main, base)

	return r
}

func setLoginRegister(r *gin.RouterGroup, repository repo.BaseRepository) {
	userHandler := InitUserApi(repository)

	r.POST("/login", userHandler.Login)
	r.POST("/register", userHandler.Register)
}

func setBookRouter(r *gin.RouterGroup, base repo.BaseRepository) {
	bookHandler := InitProductApi(base)

	book := r.Group("book", auth.TokenAuthMiddleware())
	book.GET("/findAll", bookHandler.FindAll)
	book.GET("/find/:id", bookHandler.FindByID)
	book.POST("/create", bookHandler.Create)
	book.PUT("/update/:id", bookHandler.Update)
	book.DELETE("/delete/:id", bookHandler.Delete)
}

func main() {
	base := initDB()
	r := initRouter(base)

	err := r.Run(":8010")
	if err != nil {
		panic(err)
	}
}
