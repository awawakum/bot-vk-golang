package main

import (
	"bot-vk/internal/config"
	"bot-vk/internal/database"
	"bot-vk/internal/handler"
	"bot-vk/internal/service"
	"bot-vk/internal/vk"
	"log"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("False init config: ", err)
	}

	db, err := database.ConnectToDB(&cfg.DBConfig)
	if err != nil {
		log.Fatal("False db connection: ", err)
	}

	database := database.NewDataBase(db)
	services := service.NewService(database)
	handlers := handler.NewHandler(services)

	bot, err := vk.NewBot(cfg.VkToken, handlers)
	if err != nil {
		log.Fatal("Failed to initialize Vk bot: ", err)
	}

	if err := bot.Run(); err != nil {
		log.Fatal("Failed to run the bot: ", err)
	}
}
