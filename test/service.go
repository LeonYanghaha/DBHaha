package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int
const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)
var palette = []color.Color{color.White, color.Black}


func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func handler(w http.ResponseWriter, r *http.Request) {
	lissajous(w)
	//mu.Lock()
	//count++
	//mu.Unlock()
	//_, _ = fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	//for k, v := range r.Header {
	//	_, _ = fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	//}
	//_, _ = fmt.Fprintf(w, "Host = %q\n", r.Host)
	//_, _ = fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	//if err := r.ParseForm(); err != nil {
	//	log.Print(err)
	//}
	//for k, v := range r.Form {
	//	_, _ = fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	//}
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, _ *http.Request) {
	mu.Lock()
	_, _ = fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func lissajous(out io.Writer) {
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
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	_ = gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
