# Use the official Golang image as the base image
FROM golang:1.23.3-alpine3.19

# Install mongodump tool
RUN apk add --no-cache mongodb-tools

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the application code and configuration files into the Docker image
COPY . .

# Set environment variables for backup scheduling and retention policy
ENV BACKUP_CRON="0 0 * * *"
ENV RETENTION_PERIOD="7"
ENV S3_ENDPOINT=""
ENV S3_ACCESS_KEY=""
ENV S3_SECRET_KEY=""
ENV S3_BUCKET_NAME=""
ENV S3_REGION=""
ENV MONGO_DB_NAME=""

# Define the entry point to start the application
ENTRYPOINT ["go", "run", "main.go"]
