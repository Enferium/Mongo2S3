package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	// Read environment variables
	backupCron := os.Getenv("BACKUP_CRON")
	retentionPeriod := os.Getenv("RETENTION_PERIOD")
	s3Endpoint := os.Getenv("S3_ENDPOINT")
	s3AccessKey := os.Getenv("S3_ACCESS_KEY")
	s3SecretKey := os.Getenv("S3_SECRET_KEY")
	s3BucketName := os.Getenv("S3_BUCKET_NAME")
	s3Region := os.Getenv("S3_REGION")
	mongoDBName := os.Getenv("MONGO_DB_NAME")

	// Initialize cron job
	c := cron.New()
	_, err := c.AddFunc(backupCron, func() {
		backupMongoDB(s3Endpoint, s3AccessKey, s3SecretKey, s3BucketName, s3Region, mongoDBName)
	})
	if err != nil {
		log.Fatalf("Error scheduling backup: %v", err)
	}
	c.Start()

	// Keep the application running
	select {}
}

func backupMongoDB(s3Endpoint, s3AccessKey, s3SecretKey, s3BucketName, s3Region, mongoDBName string) {
	// MongoDB backup logic using mongodump
	dumpDir := "/tmp/mongodump"
	err := os.MkdirAll(dumpDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating dump directory: %v", err)
	}

	cmd := exec.Command("mongodump", "--db", mongoDBName, "--out", dumpDir)
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Error running mongodump: %v", err)
	}

	// S3 integration
	minioClient, err := minio.New(s3Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(s3AccessKey, s3SecretKey, ""),
		Secure: true,
		Region: s3Region,
	})
	if err != nil {
		log.Fatalf("Error initializing MinIO client: %v", err)
	}

	// Upload backup to S3
	backupFile := fmt.Sprintf("%s/%s", dumpDir, mongoDBName)
	_, err = minioClient.FPutObject(context.Background(), s3BucketName, backupFile, backupFile, minio.PutObjectOptions{})
	if err != nil {
		log.Fatalf("Error uploading backup to S3: %v", err)
	}

	// Cleanup local backup files
	err = os.RemoveAll(dumpDir)
	if err != nil {
		log.Fatalf("Error cleaning up local backup files: %v", err)
	}

	fmt.Println("Backup completed successfully.")
}
