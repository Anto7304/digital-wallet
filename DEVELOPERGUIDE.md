# Development Guide

## Prerequisites

- Go 1.21+
- PostgreSQL 15+
- Redis 7+ (optional)
- Docker & Docker Compose
- Git
- Make

## Setup

### 1. Clone Repository

```bash
git clone https://github.com/Anto7304/digital-wallet.git
cd digital-wallet



#environment variables
cp .env.example .env  # Edit .env with your configuration


## common commands make dev             
make dev              # Start development server
make test             # Run all tests
make build            # Build binary
make migrate-up       # Run migrations
make docker-up        # Start Docker containers
make docker-down      # Stop Docker containers
