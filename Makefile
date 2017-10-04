.PHONY: docke-image
docker-image: http-echo
	docker build . -t http-echo:smallest

http-echo: main.go
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o $@ .

.PHONY: clean
clean:
	- rm -rf http-echo
