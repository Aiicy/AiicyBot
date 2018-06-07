package main

import (
	"os"

	"github.com/VividCortex/godaemon"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {

	f, _ := os.OpenFile("AiicyBot.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	godaemon.MakeDaemon(&godaemon.DaemonAttr{CaptureOutput: true})
	// Enable this for debug mode
	//log.SetLevel(logrus.DebugLevel)
	log.SetLevel(logrus.InfoLevel)
	log.Out = f

	StartBot()
}
