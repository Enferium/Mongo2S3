package main

import (
	"log"

	"github.com/robfig/cron/v3"
)

func StartScheduler(config *Config) {
	c := cron.New()
	c.AddFunc(config.BackupCron, func() {
		log.Println("Début de la sauvegarde MongoDB...")
		if err := BackupAndUpload(config); err != nil {
			log.Printf("Erreur de sauvegarde : %v", err)
		}
		log.Println("Nettoyage des anciennes sauvegardes...")
		nbBackupRemoved, err := CleanupOldBackups(config)
		if err != nil {
			log.Printf("Erreur de nettoyage : %v", err)
		}
		log.Printf("Nettoyage terminé, %d sauvegardes supprimées", nbBackupRemoved)
	})
	c.Start()
}
