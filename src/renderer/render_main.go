package main

import (
	"fmt"
//	"math"
	"geomath"
	"image"
	"image/color"
	"image/png"
	"draw"
	"os"
	"workspace"
)

func main() {
	s := &workspace.Scene{Name: "SceneA"}
	a := workspace.NewCube(geomath.Vector{-1, -1, -7}, 4.22, 3.89, 4)
	s.Add(a)
	Render(s)
}

func Render(s *workspace.Scene) {
	size := image.Point{256, 128}
	rect := image.Rectangle{Max: size}
	img := image.NewRGBA(rect)
	//The camera is located at (0, 0, 0) and point at (0, 0, 1).
	//The canvas is located at (0, 0, 1) from the camera origin and
	//its size is 2 by 2
	
	drawP := func(p image.Point) {
		img.Set(p.X, p.Y, color.White)
	}
	
	for i := 0; i < size.X; i++ {
		for j := 0; j < size.Y; j++ {
			img.Set(i, j, color.Black)
		}
	}

	for _, o := range(s.Objects) {
		/*for _, v := range(o.Vertices) {
			gv := geomath.Sum(o.Origin, v)
			x, y := projPointCanvas(&gv)
			fmt.Println(x, y)
			img.Set(int(float32(size.X) * x), int(float32(size.Y) * y), color.White)
		}*/
		tn := len(o.Triangles)
		var pt [3]image.Point
		for t := 0; t < tn ; t += 3  {
			for pi := 0; pi < 3; pi++ {
				gv := geomath.Sum(o.Origin, o.Vertices[o.Triangles[t + pi]])
				tx, ty := projPointCanvas(gv)
				pt[pi] = image.Pt(int(float32(size.X) * tx), int(float32(size.Y) * ty))
			}
			draw.PlotBline(pt[0], pt[1], drawP)
			draw.PlotBline(pt[1], pt[2], drawP)
			draw.PlotBline(pt[2], pt[0], drawP)
		}
	}
	//plotBLine(image.Pt(10, 100), image.Pt(200, 42), img)
	
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

func projPointCanvas(p geomath.Vector) (x, y float32) {
	x = (2 + float32(p.X / -p.Z)) / 4
	y =	(1 + float32(p.Y / p.Z)) / 2
	return
}

//Bresenham's line algorithm
/*func plotBLine(a, b image.Point, img *image.RGBA) {
	deltaX := float64(b.X - a.X)
	if deltaX == 0 {
		fmt.Println("can't draw vertical line!")
		return
	}
	deltaY := float64(b.Y - a.Y)
	errM := -1.0
	errDelta := deltaY / deltaX
	y := a.Y
	
	for x := a.X; x < b.X; x++ {
		img.Set(x, y, color.White)
		if errM = errM + errDelta ; errM > 0 {
			y++
			errM = -1.0
		}
	}
	
	//==================================
	
	deltaX := b.X - a.X
	if deltaX < 0 {
		deltaX = -deltaX
	}
	deltaY := b.Y - a.Y
	if deltaY < 0 {
		deltaY = -deltaY
	}
	var sx, sy int
	if a.X < b.X {
		sx = 1
	} else {
		sx = -1
	}
	if a.Y < b.Y {
		sy = 1
	} else {
		sy = -1
	}
	
	errM := deltaX - deltaY
	x, y := a.X, a.Y
	
	for {
		img.Set(x, y, color.White)
		if x == b.X && y == b.Y {
			break
		}
		e2 := 2 * errM
		if e2 > -deltaY {
			errM -= deltaY
			x += sx
		}
		if e2 < deltaX {
			errM += deltaX
			y += sy
		}
	}
}*/
