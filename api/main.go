package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"micro-states/modules/countries"
	"micro-states/settings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func initDb() {
	connStr := os.Getenv("DATABASE")
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	settings.DB = db
}

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gin.SetMode(os.Getenv("GIN_MODE"))

	initDb()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	countries.Router(r)
	r.Run(os.Getenv("PORT")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
