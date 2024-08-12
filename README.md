# Kubernetes
## Installation
Install docker, minikube and kubectl.

## Setup
```bash
make docker-build
minikube start --nodes 4
minikube addons enable dashboard
minikube addons enable metrics-server
minikube image load trial:latest
minikube kubectl apply -f trial-list.yaml
minikube tunnel
```
### If kubectl doesn't work with Minikube
```bash
kubectl config current-context # prints minikube if installed
kubectl config use-context minikube # sets context to minikube
```
## Additional helpful commands
```bash
minikube stop # stops cluster
minikube delete # deletes cluster
minikube node * # node related commands
minikube addons * # addons related commands
minikube kubectl * # run kubectl commands via minikube
minikube kubectl cluster-info # gets cluster information
minikube kubectl get * # gets kinds
minikube kubectl edit * # edit kinds
minikube kubectl describe * # describe kinds
minikube tunnel # exposes or tunnels services created with loadbalancer type
minikube service [name] # exposes specified service via a url
minikube dashboard # opens kubernetes dashboard
```
## Load testing
```bash
make rate-limit-test
```
