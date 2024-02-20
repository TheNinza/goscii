# Goscii

Goscii (Go-Ascii) is a simple cool cli to generate ascii art from webcam stream in your shell.

<p align="center">
  <img height="400" src="https://github.com/TheNinza/goscii/assets/62726436/470aec3c-f709-4e06-9895-6fff5ce72ba6"/>
</p>



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
