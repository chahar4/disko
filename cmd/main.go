package main

import (
	"log"

	"github.com/PatrochR/disko/db"
	"github.com/PatrochR/disko/internal/channel"
	"github.com/PatrochR/disko/internal/guild"
	"github.com/PatrochR/disko/internal/user"
	"github.com/PatrochR/disko/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
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

	//Guild Injection
	guildRepo := guild.NewRepository(db.GetDB())
	guildService := guild.NewService(guildRepo)
	guildHandler := guild.NewHandler(guildService)

	//Channel Injection
	channelRepo := channel.NewRepository(db.GetDB())
	channelService := channel.NewService(channelRepo)
	channelHandler := channel.NewHandler(channelService)

	router.InitRouter(userHandler, guildHandler, channelHandler)
	if err := router.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
