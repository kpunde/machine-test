package main

import (
	"log"
	"machine_test/handler"
	logg "machine_test/utlities/logger"
	"os"
	"os/signal"
	"syscall"
)

func cleanup() {
	log.Printf("Shutting down ...")
}

func handleInterrupt() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()
}

func main() {
	handleInterrupt()

	logger := logg.InitService("/resources/app.log", 4)
	defer func(logger logg.LoggerService) {
		err := logger.StopLog()
		if err != nil {
			log.Println(err)
		}
	}(logger)

	logger.Info("Application Started !")

	pHandler := handler.NewPortHandler("/resources/")
	pHandler.HandleFSFiles()
	pHandler.OutputAllEntries()
}
