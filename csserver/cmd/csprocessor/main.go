package main

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Processor started")

	go forever()
	select {}
}

func forever() {
	for {
		fmt.Printf("%v+\n", time.Now())
		time.Sleep(time.Second * 10)
	}
}
