## Overview
This project implements a data generation and processing pipeline using Golang and Kafka.

## Features
- Generate CSV data (id, name, address, continent)
- Produce data to Kafka topic: `source`
- Consume and sort data by:
  - ID (numerical)
  - Name (alphabetical)
  - Continent (alphabetical)
- Chunk-based sorting for memory efficiency
- Dockerized setup

## Architecture
Producer -> Kafka -> Sorter -> Chunk Files

## Run Instructions

### Start services
docker-compose up -d

### Wait for Kafka (30 sec)
sleep 30

### Create topic
docker exec -it my_project-kafka-1 kafka-topics \
--create \
--topic source \
--bootstrap-server kafka:9092 \
--partitions 1 \
--replication-factor 1

### Run producer
docker exec -it my_project-app-1 ./producer

### Run sorter
docker exec -it my_project-app-1 ./sorter

## Output
Sorted chunk files:
- chunk_id_*.txt
- chunk_name_*.txt
- chunk_continent_*.txt

## Performance
- Handles large datasets using chunk-based sorting
- Memory efficient (<2GB)
- Parallelizable design

## How to run
docker-compose build --no-cache 
docker-compose up -d
docker exec -it my_project-app-1 ./producer
docker exec -it my_project-app-1 ./sorter
docker exec -it my_project-app-1 ls
docker exec -it my_project-app-1 head chunk_id_0.txt
