package main

import (
	"github.com/a-skua/busybox-go/log"
)

func main() {
	log.Emergency("hello, world")
	log.Alert("hello, world")
	log.Critical("hello, world")
	log.Error("hello, world")
	log.Warning("hello, world")
	log.Notice("hello, world")
	log.Info("hello, world")
	log.Debug("hello, world")
}
