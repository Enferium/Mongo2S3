package main

import (
	"log"
)

func main() {
	log.Println("Starting...")
	config, err := LoadConfig()
	log.Println("Configuration loaded")
	if err != nil {
		log.Fatalf("Error on configuration load : %v", err)
		return
	}
	StartScheduler(config)
	log.Println("scheduler started")
	select {}
}
