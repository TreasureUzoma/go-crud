# Go CRUD API Starter

A simple CRUD REST API starter built with [Go](https://go.dev/), [Fiber](https://gofiber.io/), [GORM](https://gorm.io/), and PostgreSQL.

## Features

- Create, read, update, and delete posts
- Uses GORM for database ORM
- Fiber for fast, expressive HTTP routing
- PostgreSQL as the database
- JSON-based API responses with proper error handling

---

## Installation

### Prerequisites

- [Go](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Git](https://git-scm.com/downloads)

### Clone the repository

```bash
git clone https://github.com/teasureuzoma/go-crud.git
cd go-crud
```

### Install dependencies

```bash
go mod tidy
```

---

## Configuration

Create a `.env` file in the project root:

```env
DATABASE_URL=postgres://user:password@localhost:5432/go_crud?sslmode=disable
```

Make sure your PostgreSQL server is running and you’ve created the database (`go_crud` in this example).

---

## Running the app

```bash
go run main.go
```

Server will start at:

```
http://localhost:3000
```

---

## API Endpoints

### Posts

#### Create a Post

```http
POST /api/v1/posts
Content-Type: application/json
```

Body:

```json
{
  "title": "Learn Go",
  "content": "Once upon a time...",
  "author": "Leowave"
}
```

---

#### Get All Posts

```http
GET /api/v1/posts
```

Response:

```json
[
  {
    "id": 1,
    "title": "Learn Go",
    "content": "Once upon a time...",
    "author": "Leowave"
  }
]
```

---

#### Update a Post

```http
PATCH /api/v1/posts/:id
Content-Type: application/json
```

Body:

```json
{
  "title": "Learn Go Updated",
  "content": "New content",
  "author": "Leowave"
}
```

---

#### Delete a Post

```http
DELETE /api/v1/posts/:id
```

---

## Development Tips

- Use [Air](https://github.com/cosmtrek/air) (opional) for live reloading:

  ```bash
  go install github.com/cosmtrek/air@latest
  air
  ```

---

## License

MIT baby
