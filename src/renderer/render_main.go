package main

import (
	"fmt"
	"geomath"
	"image"
	"image/color"
	"image/png"
	"os"
	"workspace"
)

func main() {
	s := &workspace.Scene{Name: "SceneA"}
	a := workspace.NewCube(geomath.Vector{0, 0, -4}, 4, 4, 4)
	s.Add(a)
	Render(s)
}

func Render(s *workspace.Scene) {
	rect := image.Rect(0, 0, 128, 128)
	img := image.NewRGBA(rect)

	fmt.Println(s)
	SaveImage(img)
}

func SaveImage(img *image.RGBA) {
	f, err := os.OpenFile("../../temp/test.png", os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Image saved!")
}
