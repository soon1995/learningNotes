package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

func main() {
	// Lissajous()
	// Exercise1_5()
	Exercise1_6()
}

const (
	whiteIndex = iota // first color in palette
	blackIndex
	greenIndex
  redIndex
)

var pallette = []color.Color{color.White, color.Black,color.RGBA{0x0, 0xff, 0x0, 0xff}, color.RGBA{0xff, 0x00, 0x0, 0xff}}

// Lissajous generates GIF animations of random Lissajous figures.
func Lissajous() {
	lissajous(os.Stdout, blackIndex)
}

func lissajous(out io.Writer, colorIndex uint8) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

// Change the Lissajous program’s color palette to green on black,for
// added authenticity. To create the web color #RRGGBB, use
// color.RGBA{0xRR,0xGG, 0xBB, 0xff}, where each pair of hexadecimal digits
// represents the intensity of the red, green, or blue component of the pixel.
func Exercise1_5() {
	lissajous(os.Stdout, greenIndex)
}

func Exercise1_6() {
	lissajous(os.Stdout, redIndex)

}
