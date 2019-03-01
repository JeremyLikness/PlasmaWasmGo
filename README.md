# Plasma in Go and WebAssembly

This is a simple experiment to look at WASM performance vs. plain JavaScript. It is inspired by [this plasma effect](https://jsfiddle.net/jeremylikness/bVY6t/).

## Instructions

I assume you have [Go](https://golang.org) installed. The link will provide instructions.

If you need to run a local server, launch:

`go run goserve.go`

To compile the plasma code:

`GOOS=js GOARCH=wasm go build -o plasma.wasm plasma.go`

Then navigate to:

`http://localhost:8080`