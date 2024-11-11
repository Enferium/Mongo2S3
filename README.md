## Mongo2S3

**Mongo2S3** is a lightweight utility for automatically backing up MongoDB databases and storing backups in an S3-compatible object storage. Fully configurable through environment variables, Mongo2S3 is designed to run as a Docker container with a simple setup.

### Features
- Automated MongoDB backups using `mongodump`
- Configurable backup schedules via cron syntax
- Flexible backup retention period management
- Stores backups in any S3-compatible storage
- Environment-variable based configuration for stateless deployment

### Getting Started

#### Prerequisites
- Docker
- Access to an S3-compatible storage service (e.g., Amazon S3, MinIO)

#### Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/yourusername/mongo2s3.git
   cd mongo2s3
   ```

2. Build the Docker image:
   ```bash
   docker build -t mongo2s3 .
   ```

3. Run the container with environment variables:
   ```bash
   docker run -d --name mongo2s3-backup \
       -e MONGO_URI="mongodb://username:password@localhost:27017" \
       -e MONGO_DB="mydatabase" \
       -e S3_ENDPOINT="https://s3.amazonaws.com" \
       -e S3_ACCESS_KEY="your-access-key" \
       -e S3_SECRET_KEY="your-secret-key" \
       -e S3_BUCKET_NAME="mybackupbucket" \
       -e S3_REGION="us-west-1" \
       -e S3_PREFIX="mongo_backups/" \
       -e BACKUP_CRON="0 3 * * *" \
       -e IGNORE_COLLECTIONS="logs;temp_data" \
       -e RETENTION_DAYS=30 \
       mongo2s3
   ```

### Environment Variables

| Variable            | Description |
|---------------------|-------------|
| `MONGO_URI`         | MongoDB URI (e.g., `mongodb://username:password@localhost:27017`) |
| `MONGO_DB`          | The MongoDB database to back up |
| `S3_ENDPOINT`       | S3-compatible storage endpoint |
| `S3_ACCESS_KEY`     | S3 access key |
| `S3_SECRET_KEY`     | S3 secret key |
| `S3_BUCKET_NAME`    | S3 bucket name for backups |
| `S3_REGION`         | S3 region (default: `"auto"`) |
| `S3_PREFIX`         | Prefix for backup file storage (e.g., `"mongo_backups/"`) |
| `BACKUP_CRON`       | Cron expression for scheduling backups (e.g., `"0 3 * * *"`) |
| `IGNORE_COLLECTIONS`| Collections to exclude from backup, separated by `;` |
| `RETENTION_DAYS`    | Number of days to retain backups (e.g., `30`) |

### Usage Notes

- **Scheduling**: Use `BACKUP_CRON` to specify how often backups are created. For example, `"0 3 * * *"` will trigger a backup daily at 3:00 AM.
- **Retention**: Set `RETENTION_DAYS` to specify how long backups are kept before deletion. If `RETENTION_DAYS=0`, only the last backup is kept.
- **Collection Exclusion**: Use `IGNORE_COLLECTIONS` to avoid backing up specific collections. Collections are separated by a `;`.

### License

This project is licensed under the Apache License, Version 2.0. See the [LICENSE](LICENSE) file for details.
