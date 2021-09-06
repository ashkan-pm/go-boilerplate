# Go Boilerplate

Boilerplate for containerized Go applications including logging configuration, CI, Kubernetes resources and more

## Features
- CI pipeline using Github Actions covering lint, test and build
- Multi-stage [Dockerfile](https://github.com/ashkan-pm/go-boilerplate/blob/master/Dockerfile) optimized for cache and minimal image size
- Example Kubernetes deployment file
- Custom logger configuration supporting rolling log files and pretty logging on console using [zerolog](https://github.com/rs/zerolog)
- HTTP health-check server
- [Dotenv](https://github.com/joho/godotenv) support
- Convenient `Makefile` shell commands

## Quick Start
1. Clone the project
```shell
git clone https://github.com/ashkan-pm/go-boilerplate
```
2. Install dependencies
```shell
make install
```
3. Run the application
```shell
make run
```
4. You can `curl` the health-check to make sure everything is good
```shell
curl -v localhost:8080
```

## Docker
You can build the image by running the command below in the root of the repository
```shell
docker build . -t go-boilerplate
```

You can also run the container and expose the HTTP port by running
```shell
docker run --env ENV=production --publish 8080:8080 go-boilerplate
```

## Kubernetes
Just replace the `APP_` prefixed values in [go-boilerplate.deployment.yaml](https://github.com/ashkan-pm/go-boilerplate/blob/master/k8s/go-boilerplate.deployment.yaml) and apply on your Kubernetes cluster using
```shell
kubectl apply -f go-boilerplate.deployment.yaml
```
