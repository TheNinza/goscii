package main

import (
	"bufio"
	"image"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gocv.io/x/gocv"
)

const (
	FPS         = 30
	ImageHeight = 50
)

var asciiMap = []byte(" .:-=+*#%@")

func drawAscii(src, gray *gocv.Mat, size image.Point, buf []byte, out *bufio.Writer) ([]byte, error) {
	gocv.Resize(*src, src, size, 0, 0, gocv.InterpolationLinear)
	gocv.CvtColor(*src, gray, gocv.ColorBGRToGray)

	data, err := gray.DataPtrUint8()
	if err != nil {
		return buf, err
	}

	cols, rows := gray.Cols(), gray.Rows()

	minB, maxB := uint8(255), uint8(0)
	for _, p := range data {
		if p < minB {
			minB = p
		}
		if p > maxB {
			maxB = p
		}
	}

	rng := int(maxB) - int(minB)
	if rng == 0 {
		rng = 1
	}
	scale := len(asciiMap) - 1

	buf = buf[:0]
	buf = append(buf, 0x1b, '[', 'H') // cursor home
	for i := 0; i < rows; i++ {
		row := data[i*cols : (i+1)*cols]
		for _, p := range row {
			idx := (int(p) - int(minB)) * scale / rng
			buf = append(buf, asciiMap[idx])
		}
		buf = append(buf, '\n')
	}

	if _, err := out.Write(buf); err != nil {
		return buf, err
	}
	return buf, out.Flush()
}

func main() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		println("Error opening capture device. No camera or permission denied.")
		panic(err)
	}
	defer webcam.Close()

	aspectRatio := webcam.Get(gocv.VideoCaptureFrameWidth) / webcam.Get(gocv.VideoCaptureFrameHeight)
	size := image.Point{X: int(ImageHeight * aspectRatio), Y: ImageHeight}

	img := gocv.NewMat()
	defer img.Close()
	gray := gocv.NewMat()
	defer gray.Close()

	out := bufio.NewWriter(os.Stdout)
	buf := make([]byte, 0, (size.X+1)*size.Y+8)

	// hide cursor + clear once
	out.WriteString("\x1b[?25l\x1b[2J")
	out.Flush()

	restoreCursor := func() {
		os.Stdout.WriteString("\x1b[?25h")
	}
	defer restoreCursor()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sig
		restoreCursor()
		os.Exit(0)
	}()

	ticker := time.NewTicker(time.Second / FPS)
	defer ticker.Stop()

	for range ticker.C {
		if ok := webcam.Read(&img); !ok || img.Empty() {
			continue
		}
		buf, err = drawAscii(&img, &gray, size, buf, out)
		if err != nil {
			panic(err)
		}
	}
}
