
build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o frontend starter.go
docker:
	docker build . -t frontend:latest

