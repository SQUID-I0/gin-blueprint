# Gin Blueprint

A production-ready REST API starter kit using Gin and PostgreSQL. Skip the boilerplate and start building your API.

## What's Inside

- PostgreSQL with GORM
- Request validation that actually works
- Error handling that doesn't suck
- Rate limiting built in
- Pagination support
- CORS middleware
- Custom validators
- Soft deletes
- Ready for authentication (JWT setup included)

## Quick Start

```bash
git clone https://github.com/AliAnilKocak/gin-blueprint.git
cd gin-blueprint
cp .env.example .env
```

Edit `.env` with your database credentials:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=yourdb
```

Run it:

```bash
go mod download
go run main.go
```

Test it:

```bash
curl http://localhost:8080/health
```

## Project Structure

```
.
├── database/       # DB connection and migrations
├── handlers/       # Request handlers
├── middlewares/    # Custom middleware
├── models/         # Database models
├── utils/          # Helper functions
├── validators/     # Custom validation rules
└── main.go
```

Simple and logical. Everything has its place.

## API Examples

### Create a User

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "alex",
    "email": "alex@example.com",
    "password": "SecurePass123",
    "first_name": "Alex",
    "last_name": "Johnson"
  }'
```

Response:

```json
{
  "success": true,
  "message": "User created successfully",
  "data": {
    "id": 1,
    "username": "alex",
    "email": "alex@example.com",
    "first_name": "Alex",
    "last_name": "Johnson",
    "created_at": "2024-01-15T10:30:00Z"
  }
}
```

### List Users (with pagination)

```bash
curl "http://localhost:8080/api/v1/users?page=1&page_size=10"
```

### Get User Details

```bash
curl http://localhost:8080/api/v1/users/1
```

### Create a Post

```bash
curl -X POST http://localhost:8080/api/v1/posts \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Getting Started with Go",
    "content": "Go is awesome for building APIs...",
    "published": true,
    "user_id": 1
  }'
```

## Available Endpoints

**Users**
- `GET /api/v1/users` - List users (paginated)
- `GET /api/v1/users/:id` - Get user details
- `POST /api/v1/users` - Create user
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user
- `GET /api/v1/users/:id/posts` - Get user's posts

**Posts**
- `GET /api/v1/posts` - List all posts
- `GET /api/v1/posts/:id` - Get post details
- `POST /api/v1/posts` - Create post

## Built-in Validation

Usernames must be alphanumeric with underscores, 3-50 characters.

Passwords need:
- 8+ characters
- At least one uppercase letter
- At least one lowercase letter
- At least one number

You can easily add your own validators in `validators/` folder.

## Error Responses

All errors follow the same format:

```json
{
  "success": false,
  "error": {
    "code": "NOT_FOUND",
    "message": "User not found"
  }
}
```

Possible error codes:
- `NOT_FOUND` - Resource doesn't exist
- `VALIDATION_ERROR` - Input validation failed
- `DUPLICATE_ENTRY` - Unique constraint violation
- `UNAUTHORIZED` - Auth required
- `INTERNAL_ERROR` - Something went wrong on our end

## Rate Limiting

Default is 100 requests per minute per IP. Change it in `main.go`:

```go
rateLimiter := middlewares.NewRateLimiter(100, time.Minute)
```

## Database Models

Three models are included as examples:
- **User** - Basic user with authentication fields
- **Post** - Blog post or article
- **Tag** - Many-to-many relationship example

GORM handles migrations automatically on startup. For production, you should use proper migration tools.

## Environment Variables

Copy `.env.example` to `.env` and configure:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=yourdb
DB_SSLMODE=disable

DB_MAX_IDLE_CONNS=10
DB_MAX_OPEN_CONNS=100
DB_CONN_MAX_LIFETIME=1h
```

## What's Next?

This starter kit gives you the foundation. Here's what you'll probably want to add:

- [ ] JWT authentication (scaffolding is already there)
- [ ] Redis for caching
- [ ] Email service
- [ ] File upload handling
- [ ] Background jobs
- [ ] Docker setup
- [ ] Tests

## Docker Support

Create a `Dockerfile`:

```dockerfile
FROM golang:1.21-alpine

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o server .

EXPOSE 8080
CMD ["./server"]
```

And `docker-compose.yml`:

```yaml
version: '3.8'
services:
  postgres:
    image: postgres:14
    environment:
      POSTGRES_DB: gindb
      POSTGRES_PASSWORD: password
    ports:
      - "5432:5432"

  api:
    build: .
    ports:
      - "8080:8080"
    depends_on:
      - postgres
    environment:
      DB_HOST: postgres
```

Run with: `docker-compose up`

## Contributing

Found a bug? Have an idea? PRs are welcome.

## License

MIT - do whatever you want with it.

---

If this helped you, star the repo. If you built something cool with it, let me know!
