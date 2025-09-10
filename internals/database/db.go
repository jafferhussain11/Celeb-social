package database

import (
	"fmt"
	"os"

	"github.com/jafferhussain11/celeb-social/models/friendships"
	"github.com/jafferhussain11/celeb-social/models/posts"
	"github.com/jafferhussain11/celeb-social/models/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Config() {
	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	//this is setting up GORM wrapper
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Unable to connect to database")
		panic(err)
	}

	sql, err := db.DB()
	if err != nil {
		fmt.Println("Unable to get DB from GORM")
		panic(err)
	}

	//settings for your actual *sql.DB which is from the go packages

	if err := sql.Ping(); err != nil {
		fmt.Println("Unable to ping DB")
		panic(err)
	}

	DB = db
	DB.AutoMigrate(&users.User{}, &friendships.Friendship{}, &posts.Post{})
}
