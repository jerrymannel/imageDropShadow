package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/anthonynsimon/bild/blur"
	"github.com/fogleman/gg"
)

func check(_err error) {
	if _err != nil {
		fmt.Println(_err)
		os.Exit(1)
	}
}

func main() {
	inputFile := flag.String("i", "", "Path to input file")
	borderToggle := flag.Bool("b", false, "Toogle image border")
	transparentBg := flag.Bool("t", false, "Background is transparent")

	flag.Parse()
	if *inputFile == "" {
		check(errors.New("Missing input file"))
	}
	fmt.Println("Input file  :", *inputFile)

	reader, err := os.Open(*inputFile)
	check(err)
	defer reader.Close()

	imageExtension := filepath.Ext(*inputFile)
	imageName := strings.TrimSuffix(filepath.Base(*inputFile), imageExtension)

	inputImage, _, err := image.Decode(reader)
	check(err)

	width := inputImage.Bounds().Dx()
	height := inputImage.Bounds().Dy()

	step := 75

	newWidth := width + step
	newHeight := height + step

	step = 10

	dc := gg.NewContext(newWidth, newHeight)

	if !*transparentBg {
		dc.DrawRectangle(0, 0, float64(newWidth), float64(newHeight))
		dc.SetRGB(1, 1, 1)
		dc.Fill()
	}

	dc.DrawRectangle(40, 40, float64(width-step), float64(height-step))
	dc.SetRGBA(0.1, 0.1, 0.1, 0.2)
	dc.Fill()

	dropShadow := blur.Box(dc.Image(), 20.0)

	dc = gg.NewContext(newWidth, newHeight)

	dc.DrawImage(dropShadow, 0, 0)
	dc.Fill()

	dc.DrawImage(inputImage, 25, 25)
	dc.Fill()

	if *borderToggle {
		dc.DrawRectangle(25, 25, float64(width), float64(height))
		dc.SetRGB(0.1, 0.1, 0.1)
		dc.Stroke()
	}

	dc.SavePNG(imageName + "-out" + imageExtension)

	fmt.Println("Output file :", imageName+"-out"+imageExtension)
}
