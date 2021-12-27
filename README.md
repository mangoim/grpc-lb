# Overview
gRPC load balancing code example.

## Build Go binary

### Client
```bash
GOOS=linux GOARCH=amd64 go build -o bin/grpc-client-svc main.go
```

### Server
```bash
GOOS=linux GOARCH=amd64 go build -o bin/grpc-server-svc main.go
```

## Building Docker image
### Client
```bash
docker build -t mangobeta/grpc-lb-client:1.0.1 .
docker image push mangobeta/grpc-lb-client:1.0.1
```
### Server
```bash
docker build -t mangobeta/grpc-lb-server:1.0.1 .
docker image push mangobeta/grpc-lb-server:1.0.1
```

### Kubernetes
```bash
kubectl apply -f deployment-server.yaml
kubectl apply -f service.yaml
kubectl apply -f deployment-client.yaml
```
