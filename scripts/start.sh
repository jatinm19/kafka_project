#!/bin/bash

docker-compose up -d

echo "Waiting for Kafka..."
sleep 20

docker exec -it my_project-app-1 ./producer
docker exec -it my_project-app-1 ./sorter