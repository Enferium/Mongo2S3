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
		Region: config.S3Region,
	})
	if err != nil {
		return 0, err
	}

	ctx := context.Background()
	objectCh := minioClient.ListObjects(ctx, config.S3BucketName, minio.ListObjectsOptions{
		Prefix: config.S3Prefix,
	})

	nbBackupRemoved := 0
	var latestObject minio.ObjectInfo
	var latestObjectFound bool

	for object := range objectCh {
		if object.Err != nil {
			return nbBackupRemoved, object.Err
		}

		if config.RetentionDays == 0 {
			if !latestObjectFound || object.LastModified.After(latestObject.LastModified) {
				latestObject = object
				latestObjectFound = true
			}
		} else {
			if time.Since(object.LastModified).Hours() > float64(24*config.RetentionDays) {
				err = minioClient.RemoveObject(ctx, config.S3BucketName, object.Key, minio.RemoveObjectOptions{})
				if err != nil {
					return nbBackupRemoved, fmt.Errorf("error while removing %s: %w", object.Key, err)
				}
				nbBackupRemoved++
			}
		}
	}

	if config.RetentionDays == 0 && latestObjectFound {
		objectCh = minioClient.ListObjects(ctx, config.S3BucketName, minio.ListObjectsOptions{
			Prefix: config.S3Prefix,
		})

		for object := range objectCh {
			if object.Err != nil {
				return nbBackupRemoved, object.Err
			}
			if object.Key != latestObject.Key {
				err = minioClient.RemoveObject(ctx, config.S3BucketName, object.Key, minio.RemoveObjectOptions{})
				if err != nil {
					return nbBackupRemoved, fmt.Errorf("error while removing %s: %w", object.Key, err)
				}
				nbBackupRemoved++
			}
		}
	}

	return nbBackupRemoved, nil
}
