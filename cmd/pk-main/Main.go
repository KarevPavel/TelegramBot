package main

import "telegram-bot-long-polling/external/aria2"

func main() {
	var aria = new(aria2.AriaStarter)
	aria.EnableRPC()
	aria.Logfile("log.log")
	aria.Port(9999)
	aria.LogLevel("debug")
	aria.Start()
}
