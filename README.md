# TTL Checker API

REST API сервис для проверки срока действия (TTL) SSL-сертификатов сайтов.  
Позволяет добавлять сертификаты, проверять их валидность и отслеживать срок истечения.

# Tech Stack

- Go (Golang)
- Gin Web Framework
- PostgreSQL

# Run with Docker
1. Создать файл `.env` с настройками для PostgreSQL

```.env
DB_USERNAME= postgres
DB_PASSWORD= postgres // можно любой другой
DB_HOST= db 
DB_PORT= 5432
DB_NAME= ttl-checker // можно любое другое
DB_SSLMODE= disable
```
2. Создать в корне папку `config` , а внутри папки файл `config.yaml`

```config.yaml
port: "8000"

db:
  host: "db"
  port: "5432"
  username: "postgres"
  password: "postgres"
  dbname: "ttl-checker"
  sslmode: "disable"
```
3. Запустить проект
```
docker compose up --build
```

4. Использовать api
```
http://localhost:8000/swagger/index.html
```

# Database Migrations

Миграции выполняются автоматически при запуске через контейнер `migrate`.
SQL-файлы находятся в папке: schema

# Project Structure

- **cmd** — точка входа (main.go)  
- **pkg** — бизнес-логика (handler, service, repository)  
- **migrations** — SQL миграции  
- **docs** — Swagger документация  
- **configs** — конфигурационные файлы  
- **Dockerfile**  
- **docker-compose.yml**


