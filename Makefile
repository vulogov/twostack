all: pre test
pre:
	go get github.com/lrita/cmap
	go get github.com/gammazero/deque
	go get github.com/deckarep/golang-set
	go get github.com/cstockton/go-conv
test:
	go test -v