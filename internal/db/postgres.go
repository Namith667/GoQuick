package db

import (
	"fmt"
	"log"

	//"github.com/Namith667/GoQuick/internal/db"
	"github.com/Namith667/GoQuick/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	DSN string
	DB  *gorm.DB
}

func (p *PostgresDB) Connect() (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(p.DSN), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	p.DB = db
	return db, nil
}

func NewPostgresDB(dsn string) *PostgresDB {
	return &PostgresDB{DSN: dsn}
}

func (p *PostgresDB) RunMigrations() {
	err := p.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate models %s", err)
	}

	log.Println("DB Migrations Sucess!!")
}
