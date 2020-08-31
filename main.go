package main

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/anthonynsimon/bild/blur"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/fogleman/gg"
)

func check(_err error) {
	if _err != nil {
		panic(_err)
	}
}

func main() {
	reader, err := os.Open("login.png")
	check(err)
	defer reader.Close()
	inputImage, _, err := image.Decode(reader)
	check(err)

	width := inputImage.Bounds().Dx()
	height := inputImage.Bounds().Dy()

	newWidth := width + 50
	newHeight := height + 50

	step := 10

	dc := gg.NewContext(newWidth, newHeight)

	// dc.DrawRectangle(0, 0, float64(newWidth), float64(newHeight))
	// dc.SetRGB(1, 1, 1)
	// dc.Fill()

	dc.DrawRectangle(23, 23, float64(width+step), float64(height+step))
	dc.SetRGBA(0, 0, 0, 0.5)
	dc.Fill()
	dc.SavePNG("out-" + reader.Name())

	imageToBlur, err := imgio.Open("out-" + reader.Name())
	check(err)
	result := blur.Box(imageToBlur, 20.0)
	imgio.Save("out-blur-"+reader.Name(), result, imgio.PNGEncoder())

	intermediateImage, err := gg.LoadImage("out-blur-" + reader.Name())
	check(err)

	dc = gg.NewContext(newWidth, newHeight)
	dc.DrawImage(intermediateImage, 0, 0)
	dc.Fill()
	dc.DrawImage(inputImage, 25, 25)
	dc.Fill()
	dc.SavePNG("out-2-" + reader.Name())

}
