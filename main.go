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
	go ops.GetPortEntityFromFile("resources/ports1.json", dataChannel, errorChannel)

	isError := false

	for {
		select {
		case msg, ok := <-dataChannel:
			if !ok {
				dataChannel = nil
			}
			dataBase[msg.Name] = msg.PortObj
		case err, ok := <-errorChannel:
			if !ok {
				errorChannel = nil
			} else {
				log.Println(err)
				isError = true
			}
		}

		if dataChannel == nil && errorChannel == nil {
			break
		}
	}

	if !isError {
		for k, v := range dataBase {
			log.Printf("Port: %v, Details: %v\n", k, v)

			//Time sleep included to test the cancel signal
			time.Sleep(time.Second)
		}
	}
}
