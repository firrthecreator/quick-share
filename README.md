# Quick Share

![CI Pipeline](https://github.com/firrthecreator/quick-share/actions/workflows/ci.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/firrthecreator/quick-share)](https://goreportcard.com/report/github.com/firrthecreator/quick-share)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

A lightweight, high-performance CLI tool written in Go that enables instant file sharing between devices on the same network (Wi-Fi/LAN) via HTTP. Quick Share automatically detects your local IP address, generates a QR code in your terminal, and allows seamless file transfer without any additional setup on the receiving device.

## Features

- **Instant QR Code**: Scan the QR code from your terminal to connect immediately
- **File Hosting**: Serve files from your computer to any device on the network
- **Bidirectional Transfer**: Support for both downloading and uploading files
- **Auto-Detection**: Automatically identifies the correct local IP address
- **Single Binary**: Zero external dependencies—just download and run
- **Cross-Platform**: Works seamlessly on Windows, macOS, and Linux

## Installation

### Option 1: Install via Go

```bash
go install github.com/firrthecreator/quick-share/cmd/quick-share@latest
```

### Option 2: Build from Source

```bash
git clone https://github.com/firrthecreator/quick-share.git
cd quick-share
go mod tidy
go build -o quick-share cmd/quick-share/main.go
```

## Usage

### Share Files (Download Mode)

Share the current directory with devices on your network:

```bash
./quick-share
```

A QR code will appear in your terminal. Scan it to browse and download files.

### Receive Files (Upload Mode)

Allow other devices to upload files to your computer:

```bash
./quick-share -upload
```

Scan the QR code to access the file upload form from your phone or another device.

### Advanced Options

Share a specific directory on a custom port:

```bash
./quick-share -dir "/path/to/directory" -port 3000
```

## Project Architecture

This project follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout) for scalability and maintainability:

```
quick-share/
├── cmd/
│   └── quick-share/      # Main entry point and argument parsing
├── internal/
│   ├── network/          # Network utilities (IP detection)
│   ├── server/           # HTTP handlers (upload/download logic)
│   └── ui/               # Terminal UI and QR code generation
└── go.mod                # Dependency management
```

## Development

### Running Tests

Execute unit tests with race condition detection:

```bash
go test -v -race ./...
```

### Running Linter

Maintain code quality with strict linting:

```bash
golangci-lint run
```

## Contributing

Contributions are welcome! Please follow this workflow:

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Commit your changes: `git commit -m 'feat: add amazing feature'`
4. Push to the branch: `git push origin feature/amazing-feature`
5. Open a Pull Request

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.
