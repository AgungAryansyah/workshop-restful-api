# Prerequisites

## 1. Database

This workshop uses **PostgreSQL** (recommended). MySQL is also supported but requires minor code changes (see note below).

---

### PostgreSQL

### Option A: Install Locally

**Windows**
Download and run the installer from https://www.postgresql.org/download/windows/

**Arch Linux**
```bash
sudo pacman -S postgresql
sudo systemctl start postgresql.service
```

**Ubuntu/Debian**
```bash
sudo apt install postgresql postgresql-contrib
sudo systemctl start postgresql
```

**macOS**
```bash
brew install postgresql@13
brew services start postgresql@13
```

---

### Option B: Run with Docker (Recommended for PostgreSQL)

Make sure [Docker](https://docs.docker.com/get-docker/) is installed, then:

1. Create a `.env` file:

fill in your own values:

```
DB_NAME=
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=

APP_PORT=
```

2. Create a `docker-compose.yml` file with the following content:
```yaml
services:
  postgres:
    container_name: "postgres"
    image: postgres:13-alpine
    ports:
      - "${DB_PORT}:5432"
    environment:
      - POSTGRES_DB=${DB_NAME}
      - POSTGRES_USER=${DB_USER}
      - POSTGRES_PASSWORD=${DB_PASS}
    volumes:
      - ./postgres_data:/var/lib/postgresql/data
    networks:
      - workshop

networks:
  workshop:
    name: workshop
    driver: bridge
```

3. Start the container:
```bash
docker compose up -d
```

4. Verify it's running:
```bash
docker ps
```

For example, a filled-in `.env` for the Docker setup would look like:
```
DB_NAME=foodhub
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres

APP_PORT=8080
```

---

### MySQL (Alternative)

> ⚠️ **Note:** The workshop code is written for PostgreSQL. To use MySQL you need to:
> - Replace `gorm.io/driver/postgres` with `gorm.io/driver/mysql`
> - Replace `type:uuid` tags with `type:varchar(36)` on all entities
> - Replace `gen_random_uuid()` with `UUID()` in raw SQL

**Windows**
Download and run the installer from https://dev.mysql.com/downloads/installer/

**Arch Linux**
```bash
sudo pacman -S mysql
sudo systemctl start mysqld
```

**Ubuntu/Debian**
```bash
sudo apt install mysql-server
sudo systemctl start mysql
```

**macOS**
```bash
brew install mysql
brew services start mysql
```

**Docker**
```bash
docker run -d \
  --name mysql \
  -e MYSQL_DATABASE=your_db_name \
  -e MYSQL_USER=your_db_user \
  -e MYSQL_PASSWORD=your_db_password \
  -e MYSQL_ROOT_PASSWORD=root \
  -p 3306:3306 \
  mysql:8
```

---

## 2. Postman

**Arch Linux (AUR)**
```bash
yay -S postman-bin
```

**Other OS**
Download and install from https://www.postman.com/downloads/

Once installed, you can test the API by sending requests to `http://localhost:8080` (or whichever port you configured).

> Alternatively, you can use any HTTP client you prefer (e.g. [Insomnia](https://insomnia.rest/), [curl](https://curl.se/), or the VS Code [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) extension).
