build:
	GOOS=linux GOARCH=amd64 go build -o consignment-cli
	docker build -t consignment-cli .
run:
	docker run -e MICRO_REGISTRY=mdns consignment-cli
