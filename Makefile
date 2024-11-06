test:
	go test -v
bench:
	go test -bench=./

compile_amd64:
	GOOS=linux GOARCH=amd64 go test -c -o ./test_amd64
test_amd64: compile_amd64
	qemu-x86_64 ./test_amd64 -test.v
	rm ./test_amd64
bench_amd64: compile_amd64
	qemu-x86_64 ./test_amd64 -test.bench=./
	rm ./test_amd64

compile_arm64:
	GOOS=linux GOARCH=arm64 go test -c -o ./test_arm64
test_arm64: compile_arm64
	qemu-aarch64 ./test_arm64 -test.v
	rm ./test_arm64
bench_arm64: compile_arm64
	qemu-aarch64 ./test_arm64 -test.bench=./
	rm ./test_arm64

