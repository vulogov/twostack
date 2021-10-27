all: check
pre:
	go get github.com/lrita/cmap
	go get github.com/gammazero/deque
	go get github.com/deckarep/golang-set
	go get github.com/cstockton/go-conv
	go get github.com/google/uuid
build:
	go build -v ./...
test:
	go test -v
check: pre build test