package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	todotrpt "clean-architecture/module/item/transport"
)

func main() {
	// Checking that an environment variable is present or not.
	// mysqlConnStr := "hieu:password@tcp(127.0.0.1:3306)/clean-architecture"
	mysqlConnStr, ok := os.LookupEnv("MYSQL_CONNECTION")

	if !ok {
		log.Fatalln("Missing MySQL connection string.")
	}

	// dsn := mysqlConnStr
	dsn := mysqlConnStr + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/items", todotrpt.HandleCreateItem(db))         // create item
		v1.GET("/items", todotrpt.HandleListItem(db))            // list items
		v1.GET("/items/:id", todotrpt.HandleFindAnItem(db))      // get an item by ID
		v1.PUT("/items/:id", todotrpt.HandleUpdateAnItem(db))    // edit an item by ID
		v1.DELETE("/items/:id", todotrpt.HandleDeleteAnItem(db)) // delete an item by ID
	}

	router.Run()
}
