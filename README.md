# Blog API (Golang)

A RESTful API built with Go to manage blog posts with full CRUD functionality, in-memory storage, and an additional logical endpoint for decoding numeric strings. Includes interactive Swagger UI documentation.

---

## Running the Server

```bash
go run ./cmd/server
```
---

## The API will be accessible at:
```http request
http://localhost:8080
```
---

## API Documentation (Swagger UI)

Swagger UI is available at:
```http request
http://localhost:8080/swagger/index.html
```
From there you can:
- Explore available endpoints
- Execute requests directly in the browser
- View input/output schemas

---

## Available Endpoints

| Method  | Path        | Description                                        |
|---------|-------------|----------------------------------------------------|
| GET     | /posts      | Retrieve all blog posts                            |
| GET     | /posts/{id} | 	Retrieve a single post by ID                      |
| POST    | /posts      | Create a new blog post                             |
| PUT     | /posts/{id} | Update an existing blog post                       |
| DELETE  | /posts/{id} | Delete a blog post                                 |
| POST    | /decode     | Return the number of ways to decode a digit string |

---

## Running Tests

```bash
go test ./...
```
Runs:
- Unit tests
- API integration tests
- Logic tests for the decode endpoint

---

## Swagger Setup (if regenerating docs)
```bash
swag init -g cmd/server/main.go
```