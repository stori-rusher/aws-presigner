build: pkg-linux pkg-macos-intel pkg-macos-arm pkg-windows

pkg-linux:
	GOOS=linux GOARCH=amd64 go build -o dist/linux_amd64/presign
	tar -czvf dist/linux_amd64.tar.gz -C dist/linux_amd64/ presign

pkg-macos-intel:
	GOOS=darwin GOARCH=amd64 go build -o dist/darwin_amd64/presign
	tar -czvf dist/darwin_amd64.tar.gz -C dist/darwin_amd64/ presign

pkg-macos-arm:
	GOOS=darwin GOARCH=arm64 go build -o dist/darwin_arm64/presign
	tar -czvf dist/darwin_arm64.tar.gz -C dist/darwin_arm64/ presign

pkg-windows:
	GOOS=windows GOARCH=amd64 go build -o dist/windows_amd64/presign.exe
	zip -9 -j -y dist/windows_amd64.zip dist/windows_amd64/presign.exe

clean:
	rm -rf dist
