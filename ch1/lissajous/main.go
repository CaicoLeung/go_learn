// Lissajous generates GIF animations of random Lissajous figures.

package main

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

var (
	red   = color.RGBA{R: 255, A: 1}
	green = color.RGBA{G: 255, A: 1}
	blue  = color.RGBA{B: 255, A: 1}
	pink  = color.RGBA{R: 255, G: 192, B: 203, A: 1}
	gold  = color.RGBA{R: 255, G: 215, A: 1}
)

// 调色板
var palette = []color.Color{color.White, red, green, blue, pink, color.Black, gold}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles = 5     // 完整的x振荡器转数
		res    = 0.001 // 角分辨率
		size   = 100   // 图片画布封面[-size..+size]
		frames = 64    // 动画帧数
		delay  = 8     // 帧之间的延迟，以10ms为单位
	)

	freq := rand.Float64() * 3.0 // Y振荡器的相对频率
	anim := gif.GIF{LoopCount: frames}
	phase := 0.0 // 相位差
	for i := 0; i < frames; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(i%len(palette)))
		}
		phase += 1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	_ = gif.EncodeAll(out, &anim)
}
