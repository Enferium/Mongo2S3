# MongoDB Backup Application

This repository holds an application to backup MongoDB database to any object storage S3 compatible. It backs up on a regular basis and has a retention period.

## Building and Running the Docker Container

To containerize the application using Docker, follow these steps:

1. **Create a `Dockerfile` in the root directory of the repository to define the Docker image.**
2. **Specify the base image and any necessary dependencies in the `Dockerfile`.**
3. **Copy the application code and configuration files into the Docker image.**
4. **Set environment variables for backup scheduling and retention policy in the `Dockerfile`.**
5. **Define the entry point for the Docker container to start the application.**
6. **Build the Docker image using the `docker build` command.**
7. **Run the Docker container using the `docker run` command, passing the necessary environment variables.**

### Example Commands

```sh
# Build the Docker image
docker build -t mongodb-backup-app .

# Run the Docker container
docker run -d \
  -e BACKUP_CRON="0 0 * * *" \
  -e RETENTION_PERIOD="7" \
  -e S3_ENDPOINT="your-s3-endpoint" \
  -e S3_ACCESS_KEY="your-s3-access-key" \
  -e S3_SECRET_KEY="your-s3-secret-key" \
  -e S3_BUCKET_NAME="your-s3-bucket-name" \
  -e S3_REGION="your-s3-region" \
  -e MONGO_DB_NAME="your-mongo-db-name" \
  --name mongodb-backup-app \
  mongodb-backup-app
```

## Environment Variables

To configure the backup scheduling, retention policy, and S3 configuration, set the following environment variables:

- `BACKUP_CRON`: Cron expression for backup scheduling (e.g., "0 0 * * *" for daily backups at midnight).
- `RETENTION_PERIOD`: Retention period for backups in days (e.g., "7" for 7 days).
- `S3_ENDPOINT`: The endpoint URL of the S3 compatible storage service.
- `S3_ACCESS_KEY`: The access key for the S3 compatible storage service.
- `S3_SECRET_KEY`: The secret key for the S3 compatible storage service.
- `S3_BUCKET_NAME`: The name of the bucket where the backups will be stored.
- `S3_REGION`: The region of the S3 compatible storage service (if applicable).
- `MONGO_DB_NAME`: The name of the MongoDB database to be backed up.
