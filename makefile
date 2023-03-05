BUILD_NAME=neofetcher
.DEFAULT_GOAL := run

build:
	GOARCH=amd64 GOOS=linux go build -o ./target/${BUILD_NAME}-linux main.go
	# GOARCH=amd64 GOOS=darwin go build -o ./target/${BUILD_NAME}-darwin main.go
	# GOARCH=amd64 GOOS=windows go build -o ./target/${BUILD_NAME}-windows main.go

run: build
	./target/${BUILD_NAME}-linux
	
pushToVps:
	rsync ./target/${BUILD_NAME}-linux root@athi.fun:~/bin