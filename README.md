# Go Web RPG App

This project is a web application built with Go as the backend and HTML for the frontend. It allows users to send a "hello" message from the app to the game.

## Project Structure

```
go-web-rpg-app
├── cmd
│   └── server
│       └── main.go        # Entry point of the application
├── internal
│   ├── app
│   │   └── app.go         # Application logic
│   └── game
│       └── game.go        # Game-related functionality
├── web
│   ├── static
│   │   └── style.css       # CSS styles for the frontend
│   └── templates
│       └── index.html      # Main HTML template for the frontend
├── go.mod                  # Module definition for the Go application
├── go.sum                  # Checksums for module dependencies
└── README.md               # Documentation for the project
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd go-web-rpg-app
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Run the application:**
   ```
   go run cmd/server/main.go
   ```

4. **Access the application:**
   Open your web browser and navigate to `http://localhost:8080` to interact with the app.

## Usage Guidelines

- The application provides a simple interface to send a "hello" message to the game.
- Responses from the game will be displayed on the frontend.
- You can modify the application logic in `internal/app/app.go` and the game functionality in `internal/game/game.go` as needed.

## Contributing

Contributions are welcome! Please submit a pull request or open an issue for any enhancements or bug fixes.