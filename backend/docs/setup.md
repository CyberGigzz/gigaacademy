# 🚀 Project Setup Guide

## 1️⃣ Prerequisites

Before you start, install the following:

- **Go** (v1.23.3) → [Download](https://go.dev/dl/)
- **Docker & Docker Compose** → [Install](https://docs.docker.com/get-docker/)
- **pgAdmin** (Running in Docker) → Managed via `docker-compose.yml`
- **migrate CLI** (for database migrations)
- **modd** (for automatic code reloading)

## 2️⃣ Install Dependencies

```sh
# Clone the repository
git clone git@github.com:CyberGigzz/gigaacademy.git
cd backend

```

## 3️⃣ Setup Environment Variables

Create a `.env` file in the project root:

```ini
POSTGRES_USER=test
POSTGRES_PASSWORD=yourpassword
POSTGRES_DB=yourdatabase
PGADMIN_DEFAULT_EMAIL=test@example.com
PGADMIN_DEFAULT_PASSWORD=testpassword
```

## 4️⃣ Start Database Services

```sh
docker compose up -d
```

## 5️⃣ Apply Migrations

```sh
migrate -path migrations -database "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@localhost:5432/$POSTGRES_DB?sslmode=disable" up
```

## 6️⃣ Run the Application

Start the Go server manually:

```sh
go run cmd/api/main.go
```

Or use **modd** for automatic reloading:

```sh
modd
```

## 7️⃣ Common Issues & Fixes

**Problem:** `migrate: command not found`\
**Solution:** Ensure `$HOME/go/bin` is in your `PATH`:

```sh
export PATH=$PATH:$HOME/go/bin
```

🚀 **You're all set up!**
