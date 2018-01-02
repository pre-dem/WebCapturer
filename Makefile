GOPATH=$(PWD)

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/screenshotd src/app/main.go
	cp screenshot_deploy.conf build/screenshot.conf
	cp start.sh build/start.sh

osx:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/screenshotd src/app/main.go
	cp screenshot_deploy.conf bin/screenshot.conf
	cp start.sh build/start.sh

docker: linux
	docker build -t screenshotd:latest .
