# Students API

A RESTful API built with Go for managing student records using SQLite database.

## Project Structure

```
STUDENT-API/
├── cmd/
│   └── students-api/
│       └── main.go          # Application entry point
├── config/
│   └── local.yaml           # Local configuration file
├── internal/
│   ├── config/              # Configuration handling
│   ├── http/                # HTTP server setup
│   ├── handlers/
│   │   └── student/         # Student API handlers
│   │       └── student.go
│   ├── storage/
│   │   ├── sqlite/          # SQLite implementation
│   │   │   ├── sqlite.go
│   │   │   └── storage.go
│   │   └── storage.db       # SQLite database file
│   ├── types/
│   │   └── types.go         # Data types and models
│   └── utils/               # Utility functions
├── storage/
│   └── storage.db           # Database file
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## Prerequisites

- Go 1.25.5 or higher
- Git

## Setup Instructions

### 1. Check Go Version

First, verify that Go is installed on your system:

```bash
go version
```

### 2. Clone the Repository

```bash
git clone https://github.com/susheel7783/students-api.git
cd students-api
```

Or if starting from scratch:

```bash
mkdir students-api
cd students-api
```

### 3. Initialize Go Module

Before starting any project, initialize the Go module:

```bash
go mod init github.com/susheel7783/students-api
```

### 4. Install Dependencies

Install the required packages:

```bash
# Install cleanenv for configuration management
go get -u github.com/ilyakaznacheev/cleanenv

# Install SQLite3 driver
go get github.com/mattn/go-sqlite3
```

After installing dependencies, run:

```bash
go mod tidy
```

### 5. Version Control with Git (Using VS Code)

**Initialize and commit your project using VS Code's Git extension:**

**Step 1: Initialize Git Repository**
- Open the **Source Control** panel (Git icon in left sidebar)
- Click **"Initialize Repository"**

**Step 2: Stage Files**
- Click the **"+"** icon next to files to stage them
- Or click **"+"** next to "Changes" to stage all files

**Step 3: Commit Changes**
- Write a commit message (e.g., "Initial commit")
- Click the **checkmark (✓)** to commit

**Step 4: Create GitHub Repository**
- Go to [GitHub](https://github.com) and create a new repository named `students-api`
- **Do NOT** initialize with README

**Step 5: Connect to GitHub and Push**

Open terminal and run:

```bash
# Add remote origin
git remote add origin https://github.com/susheel7783/students-api.git

# Rename branch to main
git branch -M main

# Push to GitHub
git push -u origin main
```

**Note:** It's a good practice to commit each step during development using the VS Code Git extension.

## Configuration

Create a `config/local.yaml` file with your configuration settings:

```yaml
env: "local"
storage_path: "./storage/storage.db"
http_server:
  address: "localhost:8080"
  timeout: 4s
  idle_timeout: 60s
```

## Running the Application

### Run with Configuration File

```bash
go run cmd/students-api/main.go -config config/local.yaml
```

The server will start on the configured address (default: `localhost:8080`).

## API Endpoints

### Student Operations

- **GET** `/students` - Get all students
- **GET** `/students/:id` - Get a student by ID
- **POST** `/students` - Create a new student
- **PUT** `/students/:id` - Update a student
- **DELETE** `/students/:id` - Delete a student

### Example Requests

**Create a new student:**

```bash
curl -X POST http://localhost:8080/students \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "age": 20
  }'
```

**Get all students:**

```bash
curl http://localhost:8080/students
```

**Get a specific student:**

```bash
curl http://localhost:8080/students/1
```

## Development

### Building the Application

```bash
go build -o bin/students-api cmd/students-api/main.go
```

### Running Tests

```bash
go test ./...
```

### Best Practice: Commit Each Step

After making changes, commit using VS Code Git extension:
- Stage files with **"+"** icon
- Write descriptive commit message
- Click **checkmark (✓)** to commit

## Database

This project uses SQLite3 as the database. The database file is automatically created at `./storage/storage.db` when you run the application.

## Technologies Used

- **Go** (1.25.5) - Programming language
- **SQLite3** - Embedded database
- **cleanenv** - Configuration management
- **net/http** - HTTP server and routing

## Deployment

### Build for Production

```bash
# Build for current OS
go build -o students-api cmd/students-api/main.go

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o students-api-linux cmd/students-api/main.go

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o students-api.exe cmd/students-api/main.go
```

### Run in Production

```bash
./students-api -config config/production.yaml
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License.

## Author

**Susheel Kumar**
- GitHub: [@susheel7783](https://github.com/susheel7783)

## Acknowledgments

- Go community for excellent documentation
- SQLite for the reliable embedded database
- cleanenv for simple configuration management
