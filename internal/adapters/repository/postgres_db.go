package repository

import (
	"fmt"
	"log"

	"github.com/EdwinPirajan/Unsuscribe.git/internal/config"
	"github.com/EdwinPirajan/Unsuscribe.git/internal/core/domain"
	"github.com/EdwinPirajan/Unsuscribe.git/internal/core/ports"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBRepository struct {
	db *gorm.DB
}

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.LoadConfig().Database.Host,
		config.LoadConfig().Database.User,
		config.LoadConfig().Database.Password,
		config.LoadConfig().Database.DBName,
		config.LoadConfig().Database.Port,
		"disable",
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Database connection established")

	DB.AutoMigrate(&domain.Unsuscribe{})
}

func NewDBRepository() ports.UnsubscribeRepository {
	return &DBRepository{db: DB}
}

func (r *DBRepository) SaveUnsubscribe(email string) error {
	unsubscribe := domain.Unsuscribe{Email: email}
	return r.db.Create(&unsubscribe).Error
}
