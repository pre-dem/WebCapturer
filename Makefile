GOPATH=$(PWD)

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/WebCapturer src/app/main.go
	cp WebCapturer_deploy.conf build/WebCapturer.conf
	cp start.sh build/start.sh

osx:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/WebCapturer src/app/main.go
	cp WebCapturer_deploy.conf bin/WebCapturer.conf
	cp start.sh build/start.sh

docker: linux
	docker build -t cnwangsiyu/web-capturer:v0.0.1 .
