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
minikube kubectl explain * # Shows docs for anything
```
### Add tls support for load balancer
```bash
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -sha256 -days 365 -nodes
kubectl -n kube-system create secret tls mkcert --key key.pem --cert cert.pem
echo "kube-system/mkcert" | minikube addons configure ingress
minikube addons configure ingress
minikube addons disable ingress
minikube addons enable ingress
kubectl -n ingress-nginx get deployment ingress-nginx-controller -o yaml | grep "kube-system" # verify if kube-system/mkcert is found
```
## Load testing
```bash
make rate-limit-test
```
