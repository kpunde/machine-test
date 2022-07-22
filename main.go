package main

import (
	"log"
	"machine_test/entity"
	"machine_test/ops"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func cleanup() {
	log.Printf("Shutting down ...")
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	dataBase := make(map[string]entity.Port)

	dataChannel := make(chan entity.PortEntity)
	errorChannel := make(chan error)
	go ops.GetPortEntityFromFile("resources/ports.json", dataChannel, errorChannel)

	for msg := range dataChannel {
		dataBase[msg.Name] = msg.PortObj
	}

	for err := range errorChannel {
		log.Println(err)
	}

	for k, v := range dataBase {
		log.Printf("Port: %v, Details: %v\n", k, v)
		time.Sleep(time.Second)
	}
}
