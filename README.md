# Go Svelte


## Local kubernetes deployment

```
minikube start
```

```
kubectl create -f k8s-deployment.yml
```

```
kubectl get deployments
```

```
kubectl get pods
```

```
kubectl get pods
```

```
kubectl get svc
```

```
kubectl expose deployment go-svelte --type=NodePort --name=go-svelte --target-port=4001
```

```
kubectl get svc
```

```
minikube ip
```