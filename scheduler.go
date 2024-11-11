package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/robfig/cron/v3"
)

func StartScheduler(config *Config) {
	c := cron.New()
	c.AddFunc(config.BackupCron, func() { cronCall(config) })
	c.Start()
}

func cronCall(config *Config) {
	// Create a lock file to prevent multiple cron jobs from running at the same time
	lockFile := filepath.Join(os.TempDir(), "mongobackups3.lock")
	if _, err := os.Stat(lockFile); err == nil {
		log.Println("Another cron job is already running, skipping this run")
		return
	}
	lock, err := os.Create(lockFile)
	if err != nil {
		log.Printf("Impossible to create lock file, abort backup: %v", err)
		return
	}
	defer func() {
		lock.Close()
		log.Println("Lock file closed")
		os.Remove(lockFile)
		log.Println("Lock file removed")
	}()

	log.Println("Starting mongo dump and upload to S3...")
	if err := BackupAndUpload(config); err != nil {
		log.Printf("Error while backup and upload: %v", err)
	}
	log.Println("Cleanup old backups...")
	nbBackupRemoved, err := CleanupOldBackups(config)
	if err != nil {
		log.Printf("Error while cleaning : %v", err)
	}
	log.Printf("Cleaning ended, %d Backup removed", nbBackupRemoved)
}
