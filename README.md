# TUS File Upload Server

This project is a TUS-based file upload server implemented in Go. It uses the [tusd](https://github.com/tus/tusd) library for handling resumable file uploads and integrates with the [Gin](https://github.com/gin-gonic/gin) web framework for routing.

## Features

- Resumable file uploads using the TUS protocol.
- File storage and locking mechanisms.
- Simple web interface for uploading files.

## Prerequisites

- Go 1.23.4 or later installed on your system.
- A working internet connection to fetch dependencies.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/mariohalucyn/tus.git
   cd tus
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Usage

### Running the Server

1. Start the server:
   ```bash
   go run main.go
   ```

2. The server will run on `http://localhost:8080`.

### API Endpoints

- **Upload Endpoint**: `http://localhost:8080/files/`
  - Supports TUS protocol methods like `POST`, `PATCH`, `HEAD`, and `OPTIONS`.

## Project Structure

- `controllers/tus.go`: Contains the TUS handler setup.
- `main.go`: Entry point for the application.
- `index.html`: Frontend for testing file uploads.
- `go.mod` and `go.sum`: Dependency management files.

## Dependencies

- [tusd](https://github.com/tus/tusd): TUS protocol implementation.
- [Gin](https://github.com/gin-gonic/gin): Web framework.
- [tus-js-client](https://github.com/tus/tus-js-client): JavaScript client for TUS protocol.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [TUS Protocol](https://tus.io): Resumable file upload protocol.
- [tusd](https://github.com/tus/tusd): Go server implementation of the TUS protocol.
