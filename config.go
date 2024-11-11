package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	MongoURI          string
	MongoDB           string
	S3Endpoint        string
	S3AccessKey       string
	S3SecretKey       string
	S3BucketName      string
	S3Region          string
	S3Prefix          string
	BackupCron        string
	ignoreCollections []string
	RetentionDays     int
}

func LoadConfig() (*Config, error) {
	retentionDays, err := strconv.Atoi(getEnv("RETENTION_DAYS", "7"))
	if err != nil {
		return nil, fmt.Errorf("invalid RETENTION_DAYS: %w", err)
	}

	config := &Config{
		MongoURI:          getEnv("MONGO_URI", ""), //mongodb://localhost:27017
		MongoDB:           getEnv("MONGO_DB", ""),
		S3Endpoint:        getEnv("S3_ENDPOINT", ""),
		S3AccessKey:       getEnv("S3_ACCESS_KEY", ""),
		S3SecretKey:       getEnv("S3_SECRET_KEY", ""),
		S3BucketName:      getEnv("S3_BUCKET_NAME", ""),
		S3Region:          getEnv("S3_REGION", "auto"),
		S3Prefix:          getEnv("S3_PREFIX", ""),
		BackupCron:        getEnv("BACKUP_CRON", ""),
		ignoreCollections: []string(strings.Split(getEnv("IGNORE_COLLECTIONS", "toignore1;toignore2"), ";")),
		RetentionDays:     retentionDays,
	}

	if config.S3Endpoint == "" || config.S3AccessKey == "" || config.S3SecretKey == "" || config.S3BucketName == "" ||
		config.BackupCron == "" || config.MongoURI == "" || config.MongoDB == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	return config, nil
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
