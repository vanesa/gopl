// Lissajous generates GIF animations of random Lissajous figures
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

var seaGreen = color.RGBA{0x20, 0xb2, 0xaa, 0xff}
var orange = color.RGBA{0xcd, 0x85, 0x3f, 0xff}
var indianRed = color.RGBA{0xcd, 0x5c, 0x5c, 0xff}
var palette = []color.Color{color.Black, seaGreen, orange, indianRed}

const (
  blackIndex = 0 // first color in palette
  greenIndex = 1 // next color in palette
  orangeIndex = 2 // next color in palette
  indianredIndex = 3 // next color in palette
)

func main() {
  lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
  const (
    cycles = 5 // number of complete x oscillator revolution
    res = 0.0001 // angular resolution
    size = 100 // image canvas covers [-size ...+size]
    nframes = 64 // number of animation frames
    delay = 4 // delay between frames in 10ms units
  )

  freq := rand.Float64() * 3.0 // relative frecuendy of y oscillator
  anim := gif.GIF{LoopCount: nframes}
  phase := 0.0 // phase difference
  for i := 0; i < nframes; i++ {
    rect := image.Rect(0, 0, 2*size+1, 2*size+1)
    img := image.NewPaletted(rect, palette)
    for t := 0.0; t < cycles*2*math.Pi; t += res {
      x := math.Sin(t*1.3)
      y := math.Sin(t*freq + phase)
      img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
      greenIndex)
      img.SetColorIndex(size+int(x*size+1.2), size+int(y*size+1.2),
      orangeIndex)
      img.SetColorIndex(size+int(x*size+2.0), size+int(y*size+1.0),
      indianredIndex)
    }
    phase += 0.1
    anim.Delay = append(anim.Delay, delay)
    anim.Image = append(anim.Image, img)
  }
  gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
