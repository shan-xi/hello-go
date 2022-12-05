# hello-go
golang hello world

# docker command
```
docker build -t spinliao/hello .
docker run -d -p 80:3000 spinliao/hello
```

# gke command
```
gcloud config get-value project
gcloud compute regions list
gcloud artifacts repositories create hello-go-repo --project=go-microservice-370513 --repository-format=docker --location=us-west1 --description="Hello Go Docker repository"
gcloud builds submit --tag us-west1-docker.pkg.dev/go-microservice-370513/hello-go-repo/hello-go-gke .

gcloud container clusters create-auto hello-go-gke --region us-west1
kubectl get nodes

kubectl apply -f deployment.yaml
kubectl get deployments
kubectl get pods

kubectl apply -f service.yaml
kubectl get services

gcloud container clusters delete hello-go-gke --region us-west1
gcloud artifacts docker images delete us-west1-docker.pkg.dev/go-microservice-370513/hello-go-repo/hello-go-gke
```