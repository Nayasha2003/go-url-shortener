# Go URL Shortener 🚀

A **production-grade URL shortening service** built with **Golang**, **PostgreSQL**, **Redis**, and **Docker**.  
Designed using **real-world backend architecture**, this project showcases **high-performance caching**, **reliable data persistence**, and **containerized deployment** — exactly the skills recruiters seek in backend and software engineers.

---

## ✨ Features

- 🔗 Generate short URLs with collision-safe 6-character codes  
- ⚡ Ultra-fast redirection powered by Redis caching  
- 🗄️ Reliable and persistent URL storage using PostgreSQL  
- ❤️ Built-in health check endpoint for service monitoring  
- 🐳 Fully Dockerized setup (Go API + Redis + PostgreSQL)  
- 🧩 Clean, modular, and scalable Go project architecture  

---


## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Database:** PostgreSQL
- **Cache:** Redis
- **Containerization:** Docker & Docker Compose
- **HTTP Server:** net/http
- **Version Control:** Git & GitHub

---

## 🚀 Run the Project (Docker)

Clone the repository:

```bash
git clone https://github.com/Nayasha2003/go-url-shortener.git
cd go-url-shortener
```
Build and start all services:

```bash
docker compose up --build
```
The following services will start automatically:

Go Server → http://localhost:8080

PostgreSQL Database

Redis Cache

🔌 API Endpoints
Shorten URL
```
POST /shorten
Content-Type: application/json
```
Request:


```
{
  "url": "https://example.com"
}
```
Response:

```
{
  "short_url": "aB3dE1"
}
```
Redirect URL
```
GET /r/aB3dE1
```
Redirects to:

```
https://example.com
```
Health Check
```
GET /health
```
Response:

```
OK
```
📁 Project Structure

```
go-url-shortener/
├── cmd/
│   └── server/          # Application entry point
├── internal/
│   └── storage/         # PostgreSQL + Redis logic
├── Dockerfile           # Go build configuration
├── docker-compose.yml   # Multi-container setup
├── go.mod
└── go.sum

```
⚡ Why This Project Stands Out
Uses Redis for real-world performance optimization

Proper separation of concerns (handlers, storage, config)

Docker-first development workflow

Easily extensible into analytics, auth, or rate-limiting

Production-grade backend system — not a toy project

👩‍💻 Skills Demonstrated
Backend Engineering with Go

REST API Design

Redis Caching Strategies

SQL & PostgreSQL Integration

Docker & Docker Compose

System Design Fundamentals

Production-ready Project Structure

🌍 Deployment Ready
This service can be deployed on:

AWS / EC2

Google Cloud

Azure

Render

Railway
