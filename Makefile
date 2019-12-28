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

# Replace akshay5995 with your docker username

docker-build:
	echo "Building docker image"
	docker build . -t akshay5995/go-svelte

docker-push:
	echo "Pushing image to docker registry"
	docker login
	docker push akshay5995/go-svelte:latest

run-in-minikube:
	echo "Deploying to minikube"
	kubectl apply -f deployments/redis-master.yml
	kubectl apply -f deployments/go-deployment.yml
	echo "URL for the service"
	minikube service go-svelte --url