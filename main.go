package main

import (
	"gocv.io/x/gocv"
	"image"
	"time"
)

const (
	FPS         = 30
	ImageHeight = 50
)

func drawAscii(img *gocv.Mat, asciiBrightnessMap *[]string, aspectRatio *float64) {

	//close terminal
	print("\033[H\033[2J")

	//resize image
	gocv.Resize(*img, img, image.Point{Y: ImageHeight, X: int(ImageHeight * (*aspectRatio))}, 0, 0, gocv.InterpolationLinear)

	//convert to grayscale
	gocv.CvtColor(*img, img, gocv.ColorBGRToGray)

	maxBrightness, minBrightness := 0, 255
	for i := 0; i < img.Rows(); i++ {
		for j := 0; j < img.Cols(); j++ {
			brightness := int(img.GetUCharAt(i, j))
			maxBrightness = max(maxBrightness, brightness)
			minBrightness = min(minBrightness, brightness)
		}
	}

	asciiStr := ""

	//print ascii art
	for i := 0; i < img.Rows(); i++ {

		strLine := ""

		for j := 0; j < img.Cols(); j++ {
			brightness := int(img.GetUCharAt(i, j))
			asciiIndex := int(float64(brightness-minBrightness) / float64(maxBrightness-minBrightness) * float64(len(*asciiBrightnessMap)-1))
			strLine += (*asciiBrightnessMap)[asciiIndex]
		}

		asciiStr += strLine + "\n"
	}

	// print ascii art
	print(asciiStr)
}

func main() {

	asciiBrightnessMap := []string{" ", ".", ":", "-", "=", "+", "*", "#", "%", "@"}

	// videoCapture
	webcam, err := gocv.VideoCaptureDevice(0)
	defer func(webcam *gocv.VideoCapture) {
		err := webcam.Close()
		if err != nil {
			panic(err)
		}
	}(webcam)

	if err != nil {
		println("Error opening capture device. There is no camera connected to the system or the permission is not granted.")
		panic(err)
	}

	aspectRatio := webcam.Get(gocv.VideoCaptureFrameWidth) / webcam.Get(gocv.VideoCaptureFrameHeight)

	img := gocv.NewMat()

	tickerDuration := time.Second / time.Duration(FPS)
	ticker := time.NewTicker(tickerDuration)

	for range ticker.C {
		webcam.Read(&img)
		drawAscii(&img, &asciiBrightnessMap, &aspectRatio)
	}
}
