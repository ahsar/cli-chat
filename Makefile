.PHONY: release

lint:
	golangci-lint run ./...

release:
	export GO111MODULE=on GOPROXY=https://goproxy.cn,direct
	go mod download
	CGO_ENABLED=0 GOARCH=amd64 go build -o release/cli-chat
