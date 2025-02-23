.PHONY: build run dev

build:
	@GOARCH=wasm GOOS=js go build -C cmd/tmc_pwa -o ../../web/app.wasm
	@go build -C cmd/tmc_svc -o ../../tmc_svc

start_svc: build
	@./tmc_svc

stop_svc:
	@killall -q -HUP tmc_svc || echo

dev: stop_svc start_svc
