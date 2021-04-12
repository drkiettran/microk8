# MicroK8s sample project
**Purpose**: demonstrate how to set up a microk8s environment & deploy 
a simple golang app into a default namespace


## install microk8s:
Follow instruction from here: `https://ubuntu.com/tutorials/install-a-local-kubernetes-with-microk8s#2-deploying-microk8s`

```shell
sudo snap install microk8s --classic --channel=1.19
sudo usermod -a -G -microk8s $USER
sudo chown -f -R $USER ~/.kube

microk8s status --wait-ready
microk8s enable dashboard dns registry istio
```

## Displaying all namespaces
```shell
microk8s kubectl get all --all-namespaces
```

## Starts up a dashboard:
```shell
microk8s dashboard-proxy
```

## Installing golang 1.16.3
```shell
sudo snap install go --classic
```

## setting environment vars:
```shell
export GOPATH=/home/student/go
export GOBIN=/home/student/go/bin
```

## write sample app:
Follow instruction from here: `https://medium.com/manikkothu/build-and-deploy-apps-on-microk8s-1df26d1ddd3c`

** There are errors on this website. They are minor, however.**

## Working with local docker registry
```shell
docker build . -t hello-web
docker images
docker run hello-web -p 8080:8080
```

Visit:  `http://localhost:8080/`

## Kubernetes - Import the image into MicroK8s registry

```shell
microk8s.enable registry
docker save hello-web > hello-web.tar
microk8s.ctr -n k8s.io image import hello-web.tar
microk8s.ctr -n k8s.io images ls --name~=hello-web
```

## Kubernetes - Working with the MicroK8s registry add-on

```shell
docker build . -t localhost:32000/hello-web:v3
docker push localhost:32000/hello-web
microk8s.kubectl apply -f hello-web-deployment.yaml
microk8s.kubectl get deployments
microk8s.kubectl get pods
microk8s.kubectl expose deployment hellow-deployment --type=NodePort --name=hello-service
microk8s.kubectl get svc
```

## testing out the website:
```shell
curl http://localhost:31819
```



