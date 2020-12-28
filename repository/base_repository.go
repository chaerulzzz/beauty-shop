package repository

import (
	"beauty-shop/models"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type BaseRepository struct {
	DB     *gorm.DB
	client *redis.Client
}

var (
	Model modelInterface = &BaseRepository{}
)

type modelInterface interface {
	Initialize(driver, user, password, port, host, name string) (BaseRepository, error)
}

func (s *BaseRepository) Initialize(driver, DbUser, DbPassword, DbPort, DbHost, DbName string) (BaseRepository, error) {
	var err error
	dbUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	s.DB, err = gorm.Open(driver, dbUrl)
	if err != nil {
		return *s, err
	}
	s.DB.Debug().AutoMigrate(
		&models.User{},
		&models.Book{},
	)

	dsn := "localhost:6379"
	s.client = redis.NewClient(&redis.Options{
		Addr: dsn,
	})

	_, err = s.client.Ping().Result()
	if err != nil {
		panic(err)
	}

	return *s, nil
}
