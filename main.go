/*
Simple app to take a screenshot of the whole screen.

Author: Hakuame
Date: 25 Nov 2020

*/

package main

import (
	"fmt"
	"image/png"
	"os"
	"time"

	"github.com/kbinani/screenshot"
)

/*
TODO
- Add custom dir to save
- Add custom filename
*/

func main() {
	bounds := screenshot.GetDisplayBounds(0)

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}
    loc, _ := time.LoadLocation("Local") // Use system's local time zone
    t := time.Now().In(loc)

	fileName := fmt.Sprintf("/home/samurai/Pictures/ss_%s.png", t.Format("2006-01-02_15:05:05"))
	file, _ := os.Create(fileName)
	defer file.Close()
	png.Encode(file, img)

	fmt.Printf("Screenshot saved as: %s\n", fileName)
}
