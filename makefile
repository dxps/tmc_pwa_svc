.PHONY: build run dev

build:
	@GOARCH=wasm GOOS=js go build -C cmd/tmc-pwa -o ../../web/app.wasm
	@go build -C cmd/tmc-pwa -o ../../tmc-pwa

run: build
	@./tmc-pwa

stop:
	@killall -q -HUP tmc-pwa || echo

dev: stop run
