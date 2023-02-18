.PHONY: docker, build, apply, clean

docker: build
	@docker build -t armarny/webhook-sample:latest .
build:
	@GOOS=linux go build -o bin/webhook-sample main.go
apply:
	@kubectl apply -f manifests
clean:
	@rm -rf bin
