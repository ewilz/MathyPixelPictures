package main

import (
    "fmt"
    "syscall/js"
    "bytes"
    "encoding/base64"
    "image"
    "image/png"
)

func Pic(dx, dy int) [][]uint8 {
	// allocate
	p := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		p[i] = make([]uint8, dx)
	}

	// drawing
	for y, _ := range p {
		for x, _ := range p[y] {
			p[y][x] = uint8(x*y*y*x)
			//p[y][x] = uint8((x + y)/2)
		}
	}
	return p
}

func showPic(_ js.Value, i []js.Value) interface{} {
  imageURL1 := "data:image/png;base64,"
  imageURL2 := Show(Pic)
  fullImageURL := imageURL1 + imageURL2
  js.Global().Get("document").Call("getElementById", "img").Set("src", fullImageURL)
  fmt.Println(imageURL1 + imageURL2)
  return js.ValueOf(i[0])
}

func registerCallbacks() {
  js.Global().Set("showPic", js.FuncOf(showPic))
}

func main() {
  c := make(chan struct{}, 0)
  fmt.Println("Go WebAssembly Initialized")
  registerCallbacks()
  <-c
}



////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////// Third Party Code //////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////


// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

func Show(f func(int, int) [][]uint8) string {
	const (
		dx = 256
		dy = 256
	)
	data := f(dx, dy)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}
	return ShowImage(m)
}

func ShowImage(m image.Image) string {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}
	enc := base64.StdEncoding.EncodeToString(buf.Bytes())
  return enc
}
