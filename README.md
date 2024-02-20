# Goscii

Goscii (Go-Ascii) is a simple cool cli to generate ascii art from webcam stream in your shell.

## Requirements

- gocv [https://gocv.io/]

> Note: gocv requires OpenCV 4.0.0 or later. You can install it using:

```bash
# MacOS
brew install opencv
```

```bash
# Ubuntu
sudo apt install libopencv-dev
```

## Usage

Install dependencies:

```bash
go mod download
```

Run the cli:

```bash
go run main.go
```

Or build the binary and run it:

```bash
go build -o goscii
./goscii
```