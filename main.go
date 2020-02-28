package main

import (
	"os"
	"os/signal"
	"telegram-bot-long-polling/global_services"
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
