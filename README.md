# Plasma in Go and WebAssembly

[![Build Status](https://jeremylikness.visualstudio.com/PlasmaWasmGo/_apis/build/status/JeremyLikness.PlasmaWasmGo?branchName=master)](https://jeremylikness.visualstudio.com/PlasmaWasmGo/_build/latest?definitionId=11&branchName=master)

![Deploy status](https://jeremylikness.vsrm.visualstudio.com/_apis/public/Release/badge/e3f65447-6a0c-49c8-93ee-ac16523cb7f2/1/1)

👀 [Live Demo](https://jlikme.z13.web.core.windows.net/wasm/PlasmaWasmGo)

This is a simple experiment to look at WASM performance vs. plain JavaScript. It is inspired by [this plasma effect](https://jsfiddle.net/jeremylikness/bVY6t/). After building this, I fixed a bug and improved the performance of the JavaScript-only version [here](https://jsfiddle.net/jeremylikness/1xfh3c25/). 

Read the related blog post: [Gopher meet Plasma: A WebAssembly Experiment](https://blog.jeremylikness.com/gopher-meet-plasma-a-webassembly-experiment-4048e4d3b8d7?utm_source=jeliknes&utm_campaign=plasmawasmgo&utm_medium=github).

> This repository is continuously built and deployed using free Azure Pipelines. If you're interested in how it was setup and configured to build automatically and deploy to low cost Azure Storage Static Websites, read [Deploy WebAssembly from GitHub to Azure Storage Static Websites with Azure Pipelines](https://jlik.me/fzj).

## Instructions

I assume you have [Go](https://golang.org) installed. The link will provide instructions. This was most recently built and tested with version 1.12 on a windows/amd64 device.

If you need to run a local server, launch:

`go run goserve.go`

To compile the plasma code:

`GOOS=js GOARCH=wasm go build -o plasma.wasm plasma.go`

Then navigate to:

`http://localhost:8080`
