BUILD = go build
releaseFD = release

.PHONY: release-windows release-linux release-mac run

run:
	@go run ./...

release-windows:
	@mkdir -p release
	@echo "Building for Windows..."
	GOOS=windows GOARCH=amd64 $(BUILD) -o $(releaseFD)/2048.exe

release-linux:
	@mkdir -p release
	@echo "Building for Linux..."
	GOOS=linux GOARCH=amd64 $(BUILD) -o $(releaseFD)/2048

release-mac:
	@mkdir -p release
	@echo "Building for MacOs..."
	GOOS=darwin GOARCH=arm64 $(BUILD) -o $(releaseFD)/2048_mac

release-all: release-linux release-mac release-windows