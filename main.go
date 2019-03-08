// Based on Gophercises tutorial
// https://gophercises.com/exercises/image

package graphgenerator

import (
	"math/rand"
	"os"

	svg "github.com/ajstarks/svgo"
)

func rn(n int) int { return rand.Intn(n) }

// main - generates an SVG graph based on inputs given.
func main() {
	data := []struct {
		Label string
		Value int
	}{
		{"Train", 25000},
		{"Test", 25000},
		{"Unlabeled", 50000},
	}

	// []int{25, 25, 100}
	canvas := svg.New(os.Stdout)
	width := len(data)*80 + 10
	max := 0
	height := 300
	for _, item := range data {
		if item.Value > max {
			max = item.Value
		}
	}
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height, "fill: white")
	for i, val := range data {
		percent := val.Value * (height - 50) / max
		canvas.Rect(i*60+10, (height-50)-percent, 50, percent, "fill:rgb(77,200,232)")
		canvas.Text(i*60+35, height-20, val.Label, "font-size: 14pt; fill: black; text-anchor: middle")
		canvas.Text(i*60+35, height+1, string(val.Value), "font-size: 14pt; fill: black; text-anchor: middle")

	}

	canvas.End()
}

// Note: Commented out for SVGo version
// w, h := len(data)*60+10, 100
// r := image.Rect(0, 0, w, h)
// img := image.NewRGBA(r)
// bg := image.NewUniform(color.RGBA{240, 240, 240, 255})
// // blue := image.NewUniform(color.RGBA{250, 180, 180, 255})

// draw.Draw(img, r, bg, image.Point{0, 0}, draw.Src)

// // for y := 0; y < h; y++ {
// // 	for x := 0; x < w; x++ {
// // 		img.Set(x, y, color.RGBA{255, 255, 255, 255})
// // 	}
// // }

// mask := image.NewRGBA(image.Rect(0, 0, w, h))
// for y := 0; y < h; y++ {
// 	for x := 0; x < w; x++ {
// 		alpha := uint8(0)
// 		switch {
// 		case y < 30:
// 			alpha = 255
// 		case y < 50:
// 			alpha = 100
// 		}
// 		mask.Set(x, y, color.RGBA{
// 			R: uint8((x + y) & 255),
// 			G: uint8((x + y) << 1 & 255),
// 			B: uint8((x + y) << 2 & 255),
// 			A: alpha,
// 		})
// 	}
// }

// for i, dp := range data {
// 	x0, y0 := (i*60 + 10), 100-dp
// 	x1, y1 := (i+1)*60-1, 100
// 	bar := image.Rect(x0, y0, x1, y1)
// 	grey := image.NewUniform(color.RGBA{180, 180, 180, 255})
// 	draw.Draw(img, bar, grey, image.Point{0, 0}, draw.Src)

// 	red := image.NewUniform(color.RGBA{255, 0, 0, 255})
// 	draw.DrawMask(img, bar, red, image.Point{0, 0}, mask, image.Point{x0, y0}, draw.Over)

// 	// for x := i*60 + 10; x < (i+1)*60; x++ {
// 	// 	for y := 100; y >= (100 - dp); y-- {
// 	// 		img.Set(x, y, color.RGBA{180, 180, 250, 255})
// 	// 	}
// 	// }
// }

// f, err := os.Create("image.png")
// if err != nil {
// 	panic(err)
// }
// defer f.Close()
// err = png.Encode(f, img)
// if err != nil {
// 	panic(err)
// }
// }
