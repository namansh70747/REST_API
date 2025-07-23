# REST_API-1

A simple, modular RESTful API for managing students, built with Go, SQLite, and a clean architecture approach.

## Features
- Add, retrieve, and list students
- SQLite database for persistent storage
- Clean, modular code structure (handlers, storage, config, types)
- Graceful server shutdown
- Configuration via YAML or environment variables
- Input validation and structured JSON responses

## Project Structure
```
REST_API-1/
├── cmd/students-api/         # Main entrypoint for the API server
├── config/                   # Configuration files (YAML)
├── internal/
│   ├── config/               # Config loading logic
│   ├── http/handlers/student # HTTP handlers for student endpoints
│   ├── sqlite/               # SQLite storage implementation
│   ├── storage/              # Storage interface
│   ├── types/                # Shared types (e.g., Student struct)
│   └── utils/response/       # JSON response helpers
├── storage/                  # SQLite database file
└── go.mod, go.sum            # Go module files
```

## Getting Started

### Prerequisites
- Go 1.18+
- [SQLite3](https://www.sqlite.org/download.html) (for local development)

### Installation
1. **Clone the repository:**
   ```sh
   git clone https://github.com/yourusername/REST_API-1.git
   cd REST_API-1
   ```
2. **Install dependencies:**
   ```sh
   go mod tidy
   ```
3. **Configure the app:**
   Edit `config/local.yaml` as needed:
   ```yaml
   env: dev
   storage_path: storage/storage.db
   http_server:
     address: localhost:8082
   ```

## Running the API Server
```sh
go run cmd/students-api/main.go -config config/local.yaml
```

## API Endpoints

### Create a Student
- **POST** `/api/students`
- **Body:**
  ```json
  {
    "name": "John Doe",
    "email": "john@example.com",
    "age": 21
  }
  ```
- **Response:** `201 Created`

### Get a Student by ID
- **GET** `/api/students/{id}`
- **Response:**
  ```json
  {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "age": 21,
    "createdAt": "2024-07-24T12:34:56Z",
    "updatedAt": "2024-07-24T12:34:56Z"
  }
  ```

### List All Students
- **GET** `/api/students/list`
- **Response:**
  ```json
  [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "age": 21,
      "createdAt": "2024-07-24T12:34:56Z",
      "updatedAt": "2024-07-24T12:34:56Z"
    },
    ...
  ]
  ```

## Configuration
- Configuration is loaded from a YAML file (e.g., `config/local.yaml`) or environment variables.
- You can specify the config file path with the `-config` flag or `CONFIG_PATH` environment variable.

## Graceful Shutdown
- The server supports graceful shutdown on `Ctrl+C` (SIGINT).

## Contributing
1. Fork the repo
2. Create your feature branch (`git checkout -b feature/YourFeature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin feature/YourFeature`)
5. Open a pull request

## License
MIT 