test:
	go test -v ./...

test_amd64:
	GOOS=linux GOARCH=amd64 go test -exec="qemu-x86_64" ./... -test.v

test_arm64:
	GOOS=linux GOARCH=arm64 go test -exec="qemu-aarch64" ./... -test.v


