package main

import (
    "fmt"
    "syscall/js"
    "bytes"
    "encoding/base64"
    "image"
    "image/png"
    "encoding/json"
)

func showPic(_ js.Value, i []js.Value) interface{} {
  pixelData := make([][]uint8, 256)
  json.Unmarshal([]byte(i[0].String()), &pixelData)
  imageURL1 := "data:image/png;base64,"
  imageURL2 := Show(pixelData)
  fullImageURL := imageURL1 + imageURL2
  js.Global().Get("document").Call("getElementById", "img").Set("src", fullImageURL)
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
//////////// (modified) Third Party Code ///////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////


// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

func Show(pixelData [][]uint8) string {
	const (
		dx = 256
		dy = 256
	)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := pixelData[y][x]
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
