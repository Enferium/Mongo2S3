package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func BackupAndUpload(config *Config) error {
	timestamp := time.Now().Format("20060102T150405")
	fileName := fmt.Sprintf("backup-%s.gz", timestamp)
	dumpCmd := exec.Command("mongodump", "--uri", config.MongoURI, "--archive="+fileName, "--gzip")

	// Exécution de la commande de sauvegarde
	if err := dumpCmd.Run(); err != nil {
		return fmt.Errorf("erreur lors de la sauvegarde MongoDB: %w", err)
	}

	// Initialisation du client S3
	minioClient, err := minio.New(config.S3Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.S3AccessKey, config.S3SecretKey, ""),
		Secure: true,
		Region: config.S3Region,
	})
	if err != nil {
		return fmt.Errorf("erreur de connexion à S3: %w", err)
	}

	// Upload du fichier sur S3
	_, err = minioClient.FPutObject(context.Background(), config.S3BucketName, fileName, fileName, minio.PutObjectOptions{})
	if err != nil {
		return fmt.Errorf("erreur d'upload vers S3: %w", err)
	}

	// Suppression du fichier local après upload
	return os.Remove(fileName)
}
