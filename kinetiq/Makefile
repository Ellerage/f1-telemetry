run:
	go run ./cmd/telemetry.go

build_win:
	CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o app.exe ./cmd/telemetry.go


build_mac:
	CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -o app ./cmd/telemetry.go