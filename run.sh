#!/bin/bash
cd docker
docker-compose down
docker-compose up -d

echo "Waiting postgresql server to launch on 5432..."
echo "Waiting Postgresql launched"
sleep 3

cd ..
go test -v