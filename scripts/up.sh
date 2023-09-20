#!/bin/bash

./scripts/swagger.sh

go build -o ggs ./cmd/ggs/ggs.go
mv ggs build/linux/ggs
docker build -t ggs .
docker-compose up -d