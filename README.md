# Hostwares Test - Go

A Go stdlib HTTP server for testing Go deployment on Hostwares.

## Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| APP_NAME | Display name | No |
| DATABASE_URL | PostgreSQL connection string | No |
| JWT_SECRET | JWT signing secret | Yes |
| SMTP_HOST | SMTP server host | No |
| S3_BUCKET | S3 bucket for file uploads | No |
| PORT | Server port (default: 8080) | No |

## Endpoints

- `GET /` — App info + env var status
- `GET /health` — Health check

## Deploy on Hostwares

1. Create a new site → select this repo
2. Set environment variables
3. Deploy! (compiles to single static binary — fast startup)

## Why Go?

- Single binary, no runtime dependencies
- ~5MB Docker image (Alpine-based)
- Sub-second cold start
- Built-in concurrency
