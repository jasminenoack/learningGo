// Package lissajous change package from main to allow to import
package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	// makes random numbers more random
	// Lissajous(os.Stdout)
	// Lissajous2(os.Stdout)
	// Lissajous3(os.Stdout)
	Lissajous4(os.Stdout)
	// ColorGifs(os.Stdout)
	// CheckerboardMoving(os.Stdout)
	// CheckerboardChanging(os.Stdout)
}

func stepColor(color string, goingUp bool, value uint8) {
	if goingUp {
		value++
	}
	value--
}

func randomShuffledIndexes() []uint8 {
	a := []uint8{0, 1, 2, 3, 4, 5}
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func CheckerboardChanging(out io.Writer) {
	frames := 255
	delay := 25
	offset := 0
	paletteInner := []color.Color{
		color.RGBA{255, 0x00, 0x00, 255},    // red
		color.RGBA{255, 255 / 2, 0x00, 255}, // orange
		color.RGBA{255, 255, 0x00, 255},     // yellow
		color.RGBA{0x00, 255, 0x00, 255},    // green
		color.RGBA{0x00, 0x00, 255, 255},    // blue
		color.RGBA{255, 0x00, 255, 255},     // purple
	}
	colorWidth := 50
	size := len(paletteInner) * colorWidth
	anim := gif.GIF{LoopCount: frames}
	for i := 0; i < frames; i++ {
		rect := image.Rect(0, 0, size, size)
		img := image.NewPaletted(rect, paletteInner)
		indexes := randomShuffledIndexes()
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				horBlock := ((k) / colorWidth) % len(paletteInner)
				vertBlock := ((j) / colorWidth) % len(paletteInner)
				colorIndex := (horBlock + vertBlock + offset) % len(paletteInner)
				img.SetColorIndex(j, k, indexes[colorIndex])
			}
		}
		offset++
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func CheckerboardMoving(out io.Writer) {
	frames := 255
	delay := 1
	offset := 0
	paletteInner := []color.Color{
		color.RGBA{255, 0x00, 0x00, 255},    // red
		color.RGBA{255, 255 / 2, 0x00, 255}, // orange
		color.RGBA{255, 255, 0x00, 255},     // yellow
		color.RGBA{0x00, 255, 0x00, 255},    // green
		color.RGBA{0x00, 0x00, 255, 255},    // blue
		color.RGBA{255, 0x00, 255, 255},     // purple
	}
	size := len(paletteInner) * 50
	colorWidth := size / len(paletteInner)
	anim := gif.GIF{LoopCount: frames}
	for i := 0; i < frames; i++ {
		rect := image.Rect(0, 0, size, size)
		img := image.NewPaletted(rect, paletteInner)
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				horBlock := ((k + offset) / colorWidth) % len(paletteInner)
				vertBlock := ((j + offset) / colorWidth) % len(paletteInner)
				colorIndex := (horBlock + vertBlock) % len(paletteInner)
				img.SetColorIndex(j, k, uint8(colorIndex))
			}
		}
		offset++
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func Lissajous(out io.Writer) {
	rand.Seed(time.Now().UTC().UnixNano())
	const (
		cycles  = 5
		res     = 0.0001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 10.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func Lissajous2(out io.Writer) {
	rand.Seed(time.Now().UTC().UnixNano())
	const (
		cycles  = 5
		res     = 0.0001
		size    = 100
		nframes = 64
		delay   = 8
	)
	var greenBlack = []color.Color{color.Black, color.RGBA{0, 255, 0, 255}}
	freq := rand.Float64() * 10.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, greenBlack)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func Lissajous3(out io.Writer) {
	rand.Seed(time.Now().UTC().UnixNano())
	const (
		cycles  = 5
		res     = 0.0001
		size    = 100
		nframes = 64
		delay   = 8
	)
	var palette3 = []color.Color{
		color.Black,
		color.RGBA{0, 255, 0, 255},
		color.RGBA{255, 0, 0, 255},
		color.RGBA{0, 0, 255, 255},
	}
	freq := rand.Float64() * 10.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette3)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex := i % len(palette3)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(colorIndex))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func Lissajous4(out io.Writer) {
	rand.Seed(time.Now().UTC().UnixNano())
	const (
		cycles  = 5
		res     = 0.0001
		size    = 100
		nframes = 64
		delay   = 8
	)
	var palette3 = []color.Color{
		color.Black,
		color.RGBA{0, 255, 0, 255},
		color.RGBA{255, 0, 0, 255},
		color.RGBA{0, 0, 255, 255},
	}
	freq := rand.Float64() * 10.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette3)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex := int(t) % len(palette3)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(colorIndex))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

// Circle creates a circle
type Circle struct {
	X, Y, R float64
}

// Brightness creates brightness
func (c *Circle) Brightness(x, y float64) uint8 {
	var dx, dy float64 = c.X - x, c.Y - y
	d := math.Sqrt(dx*dx+dy*dy) / c.R
	if d > 1 {
		return 0
	}
	return 255
}

func ColorGifs(out io.Writer) {
	w, h := 240, 240
	var hw, hh float64 = float64(w / 2), float64(h / 2)
	circles := []*Circle{&Circle{}, &Circle{}, &Circle{}}

	var palette = []color.Color{
		color.RGBA{0x00, 0x00, 0x00, 0xff},
		color.RGBA{0x00, 0x00, 0xff, 0xff},
		color.RGBA{0x00, 0xff, 0x00, 0xff},
		color.RGBA{0x00, 0xff, 0xff, 0xff},
		color.RGBA{0xff, 0x00, 0x00, 0xff},
		color.RGBA{0xff, 0x00, 0xff, 0xff},
		color.RGBA{0xff, 0xff, 0x00, 0xff},
		color.RGBA{0xff, 0xff, 0xff, 0xff},
	}
	var images []*image.Paletted
	var delays []int
	steps := 20

	for step := 0; step < steps; step++ {
		img := image.NewPaletted(image.Rect(0, 0, w, h), palette)
		images = append(images, img)
		delays = append(delays, 0)

		θ := 2.0 * math.Pi / float64(steps) * float64(step)
		for i, circle := range circles {
			θ0 := 2 * math.Pi / 3 * float64(i)
			circle.X = hw - 40*math.Sin(θ0) - 20*math.Sin(θ0+θ)
			circle.Y = hh - 40*math.Cos(θ0) - 20*math.Cos(θ0+θ)
			circle.R = 50
		}

		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				img.Set(x, y, color.RGBA{
					circles[0].Brightness(float64(x), float64(y)),
					circles[1].Brightness(float64(x), float64(y)),
					circles[2].Brightness(float64(x), float64(y)),
					255,
				})
			}
		}
	}
	gif.EncodeAll(out, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}
