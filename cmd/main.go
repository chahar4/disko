package main

import (
	"log"

	"github.com/PatrochR/disko/db"
	"github.com/PatrochR/disko/internal/user"
	"github.com/PatrochR/disko/router"
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

	router.InitRouter(userHandler)
	if err := router.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
