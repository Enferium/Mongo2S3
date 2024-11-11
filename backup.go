package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func BackupAndUpload(config *Config) error {
	timestamp := time.Now().Format("20060102T150405")
	fileName := fmt.Sprintf("backup-%s.gz", timestamp)
	dumpCmd := exec.Command("mongodump", "--uri", config.MongoURI, "--archive="+fileName, "--gzip", "--db", config.MongoDB)

	for _, collection := range config.ignoreCollections {
		dumpCmd.Args = append(dumpCmd.Args, "--excludeCollection="+collection)
	}

	log.Printf("mongo dump command: %v", dumpCmd)

	// Exécution de la commande de sauvegarde
	if err := dumpCmd.Run(); err != nil {
		return fmt.Errorf("error while backuping MongoDB: %w", err)
	}

	// Initialisation du client S3
	minioClient, err := minio.New(config.S3Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.S3AccessKey, config.S3SecretKey, ""),
		Secure: true,
		Region: config.S3Region,
	})
	if err != nil {
		return fmt.Errorf("error connecting to object storage: %w", err)
	}

	s3Key := filepath.Join(config.S3Prefix, fileName)

	// Upload du fichier sur S3
	_, err = minioClient.FPutObject(context.Background(), config.S3BucketName, s3Key, fileName, minio.PutObjectOptions{})
	if err != nil {
		return fmt.Errorf("upload error: %w", err)
	}

	// Suppression du fichier local après upload
	return os.Remove(fileName)
}
