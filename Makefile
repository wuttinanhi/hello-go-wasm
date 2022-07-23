build:
	export GOOS=js
	export GOARCH=wasm
	go build -o main.wasm .
