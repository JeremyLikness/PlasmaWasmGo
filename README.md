# Plasma in Go and WebAssembly

This is a simple experiment to look at WASM performance vs. plain JavaScript. It is inspired by [this plasma effect](https://jsfiddle.net/jeremylikness/bVY6t/). After building this, I fixed a bug and improved the performance of the JavaScript-only version [here](https://jsfiddle.net/jeremylikness/1xfh3c25/). 

Read the related blog post: [Gopher meet Plasma: A WebAssembly Experiment](https://blog.jeremylikness.com/gopher-meet-plasma-a-webassembly-experiment-4048e4d3b8d7?utm_source=jeliknes&utm_campaign=plasmawasmgo&utm_medium=githubb).

## Instructions

I assume you have [Go](https://golang.org) installed. The link will provide instructions. This was most recently built and tested with version 1.12 on a windows/amd64 device.

If you need to run a local server, launch:

`go run goserve.go`

To compile the plasma code:

`GOOS=js GOARCH=wasm go build -o plasma.wasm plasma.go`

Then navigate to:

`http://localhost:8080`
