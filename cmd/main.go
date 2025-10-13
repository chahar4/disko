package main

import (
	"log"

	"github.com/PatrochR/disko/db"
	"github.com/PatrochR/disko/internal/user"
	"github.com/PatrochR/disko/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil{
		log.Fatal(err)
	}
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatalln(err)
	}

	// User Injection
	userRepo := user.NewRepository(db.GetDB())
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	gin := gin.New()
	router := router.NewRouter(":8008", gin, userHandler)

	if err := router.Start(); err != nil {
		log.Fatal(err)
	}
}
