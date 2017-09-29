.PHONY: docke-image
docker-image: http-echo
	docker build . -t http-echo:smallest

http-echo: $(wildcard cmd/http-echo/**/*)
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' ./cmd/http-echo

.PHONY: clean
clean:
	- rm -rf http-echo
