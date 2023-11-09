clean:
	rm -rf ./encodingchecker
	rm -rf ./encodingchecker.exe

windowsbuild:
	GOOS='windows' GOARCH='amd64' go build .

tidy:
	go mod tidy

.PHONY: clean windowsbuild tidy