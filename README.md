# ⚡ Go URL Shortener

A production-ready, high-performance URL shortener built with **Go**, **PostgreSQL**, **Redis**, and **Docker** — featuring Prometheus metrics, Kubernetes deployment support, and a CI/CD pipeline via GitHub Actions.

---

## 🚀 Features

- 🔗 **Shorten any URL** — generates a unique 6-character short code instantly
- ⚡ **Redis caching** — lightning-fast redirects with 24-hour cache TTL
- 🗄️ **PostgreSQL persistence** — durable storage for all URL mappings
- 📊 **Prometheus metrics** — built-in observability at `/metrics`
- 🐳 **Docker Compose** — one command to run everything locally
- ☸️ **Kubernetes ready** — includes deployment manifests with liveness & readiness probes
- 🔁 **CI/CD pipeline** — automated build and push via GitHub Actions

---

## 🛠️ Tech Stack

| Layer | Technology |
|---|---|
| Language | Go 1.25 |
| Database | PostgreSQL 15 |
| Cache | Redis 7 |
| Containerization | Docker + Docker Compose |
| Orchestration | Kubernetes |
| Metrics | Prometheus |
| CI/CD | GitHub Actions |

---

## 📁 Project Structure

```
go-url-shortener/
├── cmd/
│   └── server/
│       └── main.go          # Entry point, HTTP handlers, Prometheus metrics
├── internal/
│   └── storage/
│       ├── postgres.go      # PostgreSQL connection and table setup
│       ├── redis.go         # Redis connection
│       └── url.go           # Insert and fetch URL mappings
├── k8s/
│   ├── deployment.yaml      # Kubernetes Deployment + ConfigMap + Secret
│   └── service.yaml         # Kubernetes Service
├── .github/
│   └── workflows/
│       └── ci.yml           # GitHub Actions CI/CD pipeline
├── docker-compose.yml       # Local development setup
├── Dockerfile               # Multi-stage build
├── go.mod
└── go.sum
```

---

## ⚙️ Getting Started

### Prerequisites

- [Docker](https://www.docker.com/) installed
- [Docker Compose](https://docs.docker.com/compose/) installed

### Run Locally

```bash
# Clone the repo
git clone https://github.com/Nayasha2003/go-url-shortener.git
cd go-url-shortener

# Start all services (app + PostgreSQL + Redis)
docker compose up --build
```

The server will be available at **`http://localhost:8080`**

---

## 📡 API Endpoints

### `GET /`
Returns API usage info.

```json
{
  "status": "ok",
  "shorten": "POST /shorten with {\"url\": \"https://example.com\"}",
  "redirect": "GET /r/<short_code>"
}
```

---

### `POST /shorten`
Shortens a URL and returns a short code.

**Request:**
```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://github.com/Nayasha2003"}'
```

**Response:**
```json
{
  "short_url": "aB3xYz"
}
```

---

### `GET /r/<short_code>`
Redirects to the original URL.

```bash
curl -L http://localhost:8080/r/aB3xYz
```

Responds with `HTTP 302` and redirects to the original URL.

---

### `GET /health`
Health check endpoint used by Kubernetes probes.

```
200 OK
```

---

### `GET /metrics`
Exposes Prometheus metrics.

| Metric | Description |
|---|---|
| `http_requests_total` | Total HTTP requests by method, endpoint, status |
| `redirect_latency_seconds` | Redirect latency (cache hit vs miss) |
| `cache_hits_total` | Total Redis cache hits |
| `cache_misses_total` | Total Redis cache misses |

---

## 🐳 Docker Compose Services

| Service | Port | Description |
|---|---|---|
| `app` | `8080` | Go HTTP server |
| `postgres` | `5432` | PostgreSQL database |
| `redis` | `6379` | Redis cache |

**Stop all services:**
```bash
docker compose down
```

**Stop and remove all data volumes:**
```bash
docker compose down -v
```

---

## ☸️ Kubernetes Deployment

```bash
# Apply all manifests
kubectl apply -f k8s/

# Check pods
kubectl get pods

# Check service
kubectl get svc
```

The deployment runs **2 replicas** with resource limits and liveness/readiness probes configured on `/health`.

---

## 🔁 CI/CD Pipeline

The GitHub Actions workflow (`.github/workflows/ci.yml`) automatically:

1. Builds the Docker image on every push
2. Logs into Docker Hub using repository secrets
3. Pushes the image to Docker Hub

### Required GitHub Secrets

Go to **Settings → Secrets and variables → Actions** and add:

| Secret | Value |
|---|---|
| `DOCKER_USERNAME` | Your Docker Hub username |
| `DOCKER_PASSWORD` | Your Docker Hub Personal Access Token |

---

## 🔧 Environment Variables

| Variable | Default | Description |
|---|---|---|
| `POSTGRES_DSN` | `postgres://postgres:password@localhost:5432/urlshortener?sslmode=disable` | PostgreSQL connection string |
| `REDIS_ADDR` | `localhost:6379` | Redis address |

---

## 🏗️ How It Works

```
Client → POST /shorten → Generate short code → Save to PostgreSQL + Redis → Return short code

Client → GET /r/<code> → Check Redis cache
                              ├── Cache HIT  → Redirect instantly
                              └── Cache MISS → Query PostgreSQL → Cache result → Redirect
```

---

## 👩‍💻 Author

**Nayasha** — [@Nayasha2003](https://github.com/Nayasha2003)

---

## 📄 License

This project is open source and available under the [MIT License](LICENSE).
