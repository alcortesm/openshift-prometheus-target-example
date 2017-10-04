BIN=http-echo
LOCAL_TAG=$(BIN):smallest
REMOTE_TAG=registry.starter-ca-central-1.openshift.com/alcortesm-tutorial/openshift-prometheus-target-example:latest

.PHONY: upload-docker-image
upload-docker-image: build-docker-image
	docker push $(REMOTE_TAG)

.PHONY: build-docker-image
build-docker-image: $(BIN)
	docker build . -t $(LOCAL_TAG)
	hash=$$(docker inspect $(LOCAL_TAG) --format='{{.Id}}' | cut -d':' -f2); \
		docker tag $$hash $(REMOTE_TAG)

$(BIN): main.go
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o $@ .

.PHONY: clean
clean:
	- rm -rf $(BIN)
