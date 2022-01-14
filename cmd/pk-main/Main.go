package main

import  "bitbucket.org/y4cxp543/aria2c"

func main() {
	var aria = new(aria2.AriaStarter)
	aria.EnableRPC()
	aria.Logfile("log.log")
	aria.Port(9999)
	aria.LogLevel("debug")
	aria.Start()
}
