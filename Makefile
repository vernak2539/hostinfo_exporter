.PHONY: run
run:
	go run cmd/main.go

.PHONY: build_all
build_all: build_darwin build_pi_4 build_pi_zero

.PHONY: build_darwin
build_darwin:
	env GOOS=darwin GOARCH=arm64 go build -o build/hostinfo_exporter-darwin-arm46 ./cmd/main.go

.PHONY: build_pi_4
build_pi_4:
	env GOOS=linux GOARCH=arm64 go build -o build/hostinfo_exporter-linux-arm64 ./cmd/main.go

.PHONY: build_pi_zero
build_pi_zero:
	env GOOS=linux GOARM=7 GOARCH=arm go build -o build/hostinfo_exporter-linux-armv7 ./cmd/main.go