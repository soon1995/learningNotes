// ex3.2 prints an SVG rendering of an eggbox or saddle.
package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	eggboxsvg, _ := os.OpenFile("eggbox.svg", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	Surface(eggboxsvg, eggBox)
	eggboxsvg.Close()

	saddlesvg, _ := os.OpenFile("saddle.svg", os.O_WRONLY| os.O_TRUNC|os.O_CREATE, 0666)
	Surface(saddlesvg, saddle)
	saddlesvg.Close()
}

// Surface computes an SVG rendering of a 3-D surface function.
const (
	width, height = 600, 320            //canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                //axis ranges (-xyrange ... +xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

type zfunc func(x, y float64) float64

func Surface(writer io.Writer, f zfunc) {
	fmt.Fprintf(writer, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)
			if isAnyFinite(ax, ay, bx, by, cx, cy, dx, dy) {
				continue
			}
			fmt.Fprintf(writer, "<polygon points='%g,%g %g, %g %g, %g %g, %g' />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(writer, "</svg>")
}

func corner(i, j int, f zfunc) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0.0)
	return math.Sin(r) / r
}

func isAnyFinite(fs ...float64) bool {
	for i := 0; i < len(fs); i++ {
		if isFinite(fs[i]) {
			return true
		}
	}
	return false
}

func isFinite(f float64) bool {
	if math.IsInf(f, 0) || math.IsNaN(f) {
		return true
	}
	return false
}

func eggBox(x, y float64) float64 {
	return 0.2 * (math.Cos(x) + math.Cos(y))
}

func saddle(x, y float64) float64 {
	a := 25.0
	b := 17.0
	a2 := a * a
	b2 := b * b
	return (y*y/a2 - x*x/b2)
}
