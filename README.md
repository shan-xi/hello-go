# hello-go

## Demo

    1. Using go-kit to do CRUD with MongoDB and Postgres.
    2. API: save user visit records in MongoDB and retrieve results.
    3. API: a todo list recorded in Postgres using gorm.
    4. Using docker-compose to network all things.

## Command records

    - docker command - build golang app
    ```
    docker build -t spinliao/hello .
    docker run -d -p 80:3000 spinliao/hello
    ```
    - docker command - build mongodb
    ```
    docker pull mongo:latest
    docker run -d -p 27017:27017 --name mongodb mongo:latest
    docker exec -it mongodb bash
    - mongosh
    docker logs mongodb --follow
    ```

    - gke command - deploy go app on gke
    ```
    gcloud config get-value project
    gcloud compute regions list
    gcloud artifacts repositories create hello-go-repo --project=go-microservice-370513 --repository-format=docker --location=us-west1 --description="Hello Go Docker repository"
    gcloud builds submit --tag us-west1-docker.pkg.dev/go-microservice-370513/hello-go-repo/hello-go-gke .

    gcloud container clusters create-auto hello-go-gke --region us-west1
    kubectl get nodes

    kubectl apply -f gke/deployment.yaml
    kubectl get deployments
    kubectl get pods

    kubectl apply -f gke/service.yaml
    kubectl get services

    gcloud container clusters delete hello-go-gke --region us-west1
    gcloud artifacts docker images delete us-west1-docker.pkg.dev/go-microservice-370513/hello-go-repo/hello-go-gke
    ```

    - HotReload
    ```
    github.com/githubnemo/CompileDaemon has some bugs
    After modifying mounted code, need to run docker restart CONTAINER_NAME
    ```