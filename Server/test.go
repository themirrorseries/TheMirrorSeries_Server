package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")
	t := time.NewTicker(time.Second)
	for v := range t.C { // 循环channel
		fmt.Println("hello", v)
	}

}
