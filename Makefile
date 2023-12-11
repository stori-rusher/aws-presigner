build: pkg-linux pkg-macos-intel pkg-macos-arm pkg-windows

pkg-linux:
	GOOS=linux GOARCH=amd64 go build -o dist/linux_amd64/presign

pkg-macos-intel:
	GOOS=darwin GOARCH=amd64 go build -o dist/darwin_amd64/presign

pkg-macos-arm:
	GOOS=darwin GOARCH=arm64 go build -o dist/darwin_arm64/presign

pkg-windows:
	GOOS=windows GOARCH=amd64 go build -o dist/windows_amd64/presign.exe

clean:
	rm -rf dist
