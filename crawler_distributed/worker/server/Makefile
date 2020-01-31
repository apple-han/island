
build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o work ./worker.go
docker:
	docker build . -t work:latest
