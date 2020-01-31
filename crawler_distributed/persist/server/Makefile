
build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o item ./itemsaver.go
docker:
	docker build . -t item:latest
