package main

import (
	"context"
	"fmt"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func CleanupOldBackups(config *Config) (int, error) {
	minioClient, err := minio.New(config.S3Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.S3AccessKey, config.S3SecretKey, ""),
		Secure: true,
	})
	if err != nil {
		return 0, err
	}

	ctx := context.Background()
	objectCh := minioClient.ListObjects(ctx, config.S3BucketName, minio.ListObjectsOptions{})

	nbBackupRemoved := 0

	for object := range objectCh {
		if object.Err != nil {
			return nbBackupRemoved, object.Err
		}

		// Vérifier la date pour supprimer si l'objet est trop vieux
		if time.Since(object.LastModified).Hours() > float64(24*config.RetentionDays) {
			err = minioClient.RemoveObject(ctx, config.S3BucketName, object.Key, minio.RemoveObjectOptions{})
			if err != nil {
				return nbBackupRemoved, fmt.Errorf("erreur de suppression de %s: %w", object.Key, err)
			}
			nbBackupRemoved++
		}
	}
	return nbBackupRemoved, nil
}
