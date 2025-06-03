# Fillout to Discord Service

This project is a Go service that integrates Fillout with Discord. It processes input from Fillout forms and updates Discord channels and roles based on the provided data.

## Project Structure

```
fillout-discord-service
├── cmd
│   └── main.go            # Entry point of the application
├── internal
│   ├── api
│   │   ├── endpoints.go   # HTTP endpoints for the service
│   │   └── middleware.go   # Middleware functions for request processing
│   ├── discord
│   │   └── client.go      # Discord client setup and management
│   ├── fillout
│   │   └── handler.go     # Logic for processing Fillout data
│   └── config
│       └── config.go      # Configuration settings and loading
├── go.mod                  # Module definition and dependencies
├── go.sum                  # Dependency checksums
└── README.md               # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd fillout-discord-service
   ```

2. **Install dependencies:**
   Ensure you have Go installed, then run:
   ```
   go mod tidy
   ```

3. **Configure the application:**
   Set up your environment variables or configuration files as needed. Refer to `internal/config/config.go` for the required settings.

4. **Run the application:**
   ```
   go run cmd/main.go
   ```

## Usage

- The service listens for incoming requests from Fillout. Ensure that your Fillout forms are configured to send data to the correct endpoints defined in `internal/api/endpoints.go`.
- The service will interact with Discord using the Discord API to manage roles and send messages based on the Fillout data.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for details.