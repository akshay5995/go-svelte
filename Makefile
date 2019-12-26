GOCMD=go
GOBUILD=go build
GO_FILE=main.go
BINARY_FILE=main

build-client:
	cd ./client && yarn && yarn build

build-server:
	echo "Building binary"
	${GOBUILD} -o ${BINARY_FILE} ${GO_FILE}

run:
	echo "Building binary"
	${GOCMD} run main.go

docker-build:
	echo "Building docker image"
	docker build . -t akshay5995/go-svelte
