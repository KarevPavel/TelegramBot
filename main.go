package main

import (
	"bitbucket.org/y4cxp543/telegram-bot/global_services"
	"os"
	"os/signal"
)

func main() {

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	go func() {
		<-c
		global_services.GlobalServicesStop()
		os.Exit(1)
	}()
	global_services.TelegramBot.RegisterBotCommand(global_services.CommandProcessor.ProcessSearchTorrents, "processSearchTorrents")
	global_services.TelegramBot.RegisterBotCommand(global_services.CommandProcessor.ProcessDocument, "processDocument")
	global_services.TelegramBot.RegisterBotCommand(global_services.CommandProcessor.ProcessMagnetLink, "processMagnetLink")
	global_services.TelegramBot.Start()
}