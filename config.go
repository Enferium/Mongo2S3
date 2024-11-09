package main

import (
	"os"
	"strconv"
)

type Config struct {
	MongoURI      string
	MongoDB       string
	S3Endpoint    string
	S3AccessKey   string
	S3SecretKey   string
	S3BucketName  string
	S3Region      string
	BackupCron    string
	RetentionDays int
}

func LoadConfig() (*Config, error) {
	retentionDays, err := strconv.Atoi(getEnv("RETENTION_DAYS", "0"))
	if err != nil {
		return nil, err
	}

	return &Config{
		MongoURI:      getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDB:       getEnv("MONGO_DB", "testdb"),
		S3Endpoint:    getEnv("S3_ENDPOINT", ""),
		S3AccessKey:   getEnv("S3_ACCESS_KEY", ""),
		S3SecretKey:   getEnv("S3_SECRET_KEY", ""),
		S3BucketName:  getEnv("S3_BUCKET_NAME", "backuptester"),
		S3Region:      getEnv("S3_REGION", "auto"),
		BackupCron:    getEnv("BACKUP_CRON", "00 19 * * sat"),
		RetentionDays: retentionDays,
	}, nil
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
