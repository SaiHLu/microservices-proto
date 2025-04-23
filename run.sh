#!/bin/bash

SERVICE_NAME=$1
RELEASE_VERSION=$2

# Install dependencies
sudo apt-get install -y protobuf-compiler golang-goprotobuf-dev 
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate protobuf files
protoc --go_out=./golang \
  --go_opt=paths=source_relative \
  --go-grpc_out=./golang \
  --go-grpc_opt=paths=source_relative \
  ./${SERVICE_NAME}/*.proto

# Initialize Go module
cd golang/${SERVICE_NAME}
go mod init github.com/SaiHLu/microservices-proto/golang/${SERVICE_NAME} || true
go mod tidy
cd ../..

# Configure Git
git config --global user.email "saisailuhlaing@gmail.com"
git config --global user.name "SaiHLu"

# Explicitly add only the protobuf files
git add golang/${SERVICE_NAME}/

# Create commit (only if there are changes to proto files)
if git diff-index --quiet HEAD --; then
  echo "No changes to commit"
else
  git commit -m "proto update for ${SERVICE_NAME} ${RELEASE_VERSION}"
fi

# Tag and push
git tag -fa "golang/${SERVICE_NAME}/${RELEASE_VERSION}" \
  -m "golang/${SERVICE_NAME}/${RELEASE_VERSION}"
git push origin "refs/tags/golang/${SERVICE_NAME}/${RELEASE_VERSION}"