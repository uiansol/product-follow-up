package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/uiansol/product-follow-up/internal/adapters/db/mysql"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/uiansol/product-follow-up/internal/adapters/api"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		if args[0] == "migrate" {
			migrate()
		}
	} else {
		api.RunServer()
	}
}

func migrate() {
	fmt.Println("Migrating MySQL Models")

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(gormmysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASS"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DB"))), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&mysql.ProductDB{})

	fmt.Println("Migration finished")
}
