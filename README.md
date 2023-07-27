# api-go

## Description

A simple API written in Go.

## Requirements

- Docker
- sqlx-cli
- Go 1.20

## Usage

```bash
docker-compose up --build && \
sqlx migrate run && \
go run main.go
```
