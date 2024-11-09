package main

import (
	"log"
)

func main() {
	log.Println("Démarrage du programme...")
	config, err := LoadConfig()
	log.Println("Configuration chargée")
	if err != nil {
		log.Fatalf("Erreur de chargement de la configuration: %v", err)
	}
	StartScheduler(config)
	log.Println("scjheduler démarré")
	select {} // Bloque l'exécution pour garder le programme actif
}
