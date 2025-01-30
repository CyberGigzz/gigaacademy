# 🗄 Database Documentation

## 1️⃣ Database: PostgreSQL
We use **PostgreSQL** for scalability and reliability.

## 2️⃣ Running PostgreSQL in Docker
Start the database with Docker:
```sh
docker compose up -d
```
Check running containers:
```sh
docker ps
```

## 3️⃣ Migrations with `migrate`
Create a migration:
```sh
migrate create -ext sql -dir migrations -seq init_schema
```
Apply migrations:
```sh
migrate -database "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@localhost:5432/$POSTGRES_DB?sslmode=disable" -path migrations up
```
Rollback migrations:
```sh
migrate -database "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@localhost:5432/$POSTGRES_DB?sslmode=disable" -path migrations down 1
```

Version migrations:
```sh
migrate -database "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@localhost:5432/$POSTGRES_DB?sslmode=disable" -path migrations version
```

## 4️⃣ Verifying Setup
Check if tables exist:
```sql
SELECT * FROM pg_tables WHERE schemaname = 'public';
```

## 5️⃣ Troubleshooting
🔹 **Database connection refused?** Ensure Docker is running:
```sh
docker compose up -d
```
🔹 **Table not found?** Run migrations again:
```sh
migrate up
```

🚀 **PostgreSQL is now ready!**

