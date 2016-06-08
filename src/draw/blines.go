//Bresenham's algoritms
package draw

import "image"

func PlotBline(a, b image.Point, plot func(p image.Point)) {
    /*deltaX := float64(b.X - a.X)
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
	}*/
	
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
		//img.Set(x, y, color.White)
		plot(image.Pt(x, y))
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
}