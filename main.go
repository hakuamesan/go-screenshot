package main

import (
	"fmt"
	"image/png"
	"os"
	"time"

	"github.com/kbinani/screenshot"
)

func main() {
	bounds := screenshot.GetDisplayBounds(0)

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}

	t := time.Now()

	fileName := fmt.Sprintf("/home/samurai/Pictures/ss_%s.png", t.Format("2006-01-01_15:05:05"))
	file, _ := os.Create(fileName)
	defer file.Close()
	png.Encode(file, img)

	fmt.Printf("screenshot: %v \"%s\"\n", bounds, fileName)
}
