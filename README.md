# Go Svelte Kubernetes


Dabbling with Go, Svelte, Kubernetes the _buzzword_ trio. A simple site to show the user information from github along with the public repos using the github API.


Looks like this: 

![Example not available](https://raw.githubusercontent.com/akshay5995/go-svelte/master/docs/example.png)

Prerequisites:

1. Node LTS
2. Go
3. Minikube (Local kubernetes cluster)
4. Docker Desktop
5. kubectl


Configurable via `.env` file. You can rename `.env.example` to `.env` to confiure the site with the following values:


DO NOT CHANGE `REDIS_HOST` while deploying to kubernetes.

```sh

GITHUB_TOKEN=<GITHUB_TOKEN> 
GITHUB_USER=akshay5995
SITE_NAME=Akshay Ram Vignesh
DEV_BLOG_SITE=https://akshayramvignesh.dev
REDIS_HOST=redis-master


```

`GITHUB_TOKEN` is a personal access token to get the user details and repo list.


Run using `Makefile`


```sh

make docker-build

```

```sh

make docker-push

```

```sh

make run-in-minikube

```