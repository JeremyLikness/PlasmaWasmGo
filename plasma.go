package main

import (
	"fmt"
	"math"
	"syscall/js"
)

var (
	window, doc, body, canvas, drawCtx     js.Value
	sine                                   [512]int
	color                                  [256]string
	pos1, pos3, tpos1, tpos2, tpos3, tpos4 uint16
)

const w int = 320
const h int = 200

func main() {

	setup()

	plasmaLoop := make(chan bool)

	var renderer js.Func

	renderer = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		updatePlasma()
		window.Call("setTimeout", renderer)
		return nil
	})

	window.Call("setTimeout", renderer)

	<-plasmaLoop
}

func updatePlasma() {
	tpos4 = 0
	tpos3 = pos3
	idx := 0
	for idx < h {
		tpos1 = pos1 + 5
		tpos2 = 3
		tpos3 &= 511
		tpos4 &= 511
		var jdx = 0
		for jdx < w {
			tpos1 &= 511
			tpos2 &= 511
			x := sine[tpos1] + sine[tpos2] + sine[tpos3] + sine[tpos4]
			pidx := (128 + uint8(x>>4))
			drawCtx.Set("fillStyle", color[pidx])
			drawCtx.Call("fillRect", jdx, idx, jdx+1, idx+1)
			tpos1 += 5
			tpos2 += 3
			jdx++
		}
		tpos4 += 3
		tpos3++
		idx++
	}
	pos1 += 9
	pos3 += 8
}

func createSineTable() {
	idx := 0
	for idx < 512 {
		rad := (float64(idx) * 0.703125) * 0.0174532
		sine[idx] = int(math.Sin(rad) * 1024)
		idx++
	}
}

func createPalette() {
	idx := 0
	for idx < 64 {
		r := idx << 2
		g := 255 - ((idx << 2) + 1)
		color[idx] = "rgb(" + fmt.Sprint(r) + "," + fmt.Sprint(g) + ",0)"
		g = (idx << 2) + 1
		color[idx+64] = "rgb(255," + fmt.Sprint(g) + ",0)"
		r = 255 - ((idx << 2) + 1)
		g = 255 - ((idx << 2) + 1)
		color[idx+128] = "rgb(" + fmt.Sprint(r) + "," + fmt.Sprint(g) + ",0)"
		g = (idx << 2) + 1
		color[idx+192] = "rgb(0," + fmt.Sprint(g) + ",0)"
		idx++
	}
}

func setup() {

	createSineTable()
	createPalette()

	window = js.Global()
	doc = window.Get("document")
	body = doc.Get("body")

	canvas = doc.Call("createElement", "canvas")
	canvas.Set("height", h)
	canvas.Set("width", w)
	body.Call("appendChild", canvas)

	drawCtx = canvas.Call("getContext", "2d")
}
