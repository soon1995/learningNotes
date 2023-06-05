package main

import (
	"fmt"
	"io"
	"math"
	"sync"
)

const (
	width, height = 600, 320            //canvas size in pixels
	cells         = 2200                // number of grid cells
	xyrange       = 30.0                //axis ranges (-xyrange ... +xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30)
)

const workers = 1100

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

// func SurfaceOptimized(writer io.Writer) {
//   if workers > cells >> 1 {
//     log.Println("worker shall not be more than cells")
//     return
//   }
// 	fmt.Fprintf(writer, "<svg xmlns='http://www.w3.org/2000/svg' "+
// 		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
// 		"width='%d' height='%d'>", width, height)
// 	var wg sync.WaitGroup
// 	size := cells / workers
// 	x := 0
// 	add := make(chan int, workers)
// 	for w := 0; w < workers; w++ {
// 		wg.Add(1)
// 		go func(w int) {
// 			defer wg.Done()
// 			for i := (size * w); i < (w+1)*size; i++ {
// 				for j := 0; j < cells; j++ {
// 					// fmt.Println(size * w)
// 					ax, ay := corner(i+1, j)
// 					bx, by := corner(i, j)
// 					cx, cy := corner(i, j+1)
// 					dx, dy := corner(i+1, j+1)
// 					fmt.Fprintf(writer, "<polygon points='%g,%g %g, %g %g, %g %g, %g' />\n",
// 						ax, ay, bx, by, cx, cy, dx, dy)
// 					add <- 1
// 				}
// 			}
// 		}(w)
// 	}
// 	go func() {
// 		wg.Wait()
// 		close(add)
// 	}()
// 	for v := range add {
// 		x += v
// 	}
// 	fmt.Println(x)
// 	fmt.Fprintln(writer, "</svg>")
// }

func SurfaceOptimized(writer io.Writer) {
	fmt.Fprintf(writer, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	var wg sync.WaitGroup
	for i := 0; i < cells; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < cells; j++ {
				ax, ay := corner(i+1, j)
				bx, by := corner(i, j)
				cx, cy := corner(i, j+1)
				dx, dy := corner(i+1, j+1)
				fmt.Fprintf(writer, "<polygon points='%g,%g %g, %g %g, %g %g, %g' />\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}(i)
	}
	wg.Wait()
	fmt.Fprintln(writer, "</svg>")
}

func SurfaceOptimized2(writer io.Writer) {
	fmt.Fprintf(writer, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	ch := make(chan struct{}, cells*cells)
	for i := 0; i < cells; i++ {
		go func(i int) {
			for j := 0; j < cells; j++ {
				ax, ay := corner(i+1, j)
				bx, by := corner(i, j)
				cx, cy := corner(i, j+1)
				dx, dy := corner(i+1, j+1)
				fmt.Fprintf(writer, "<polygon points='%g,%g %g, %g %g, %g %g, %g' />\n",
					ax, ay, bx, by, cx, cy, dx, dy)
				ch <- struct{}{}
			}
		}(i)
	}
	for i := 0; i < cells*cells; i++ {
		<-ch
	}
	fmt.Fprintln(writer, "</svg>")
}

func SurfaceNoOptimized(writer io.Writer) {
	fmt.Fprintf(writer, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(writer, "<polygon points='%g,%g %g, %g %g, %g %g, %g' />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(writer, "</svg>")
}

func corner(i, j int) (float64, float64) {
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
