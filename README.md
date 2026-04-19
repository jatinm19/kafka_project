# Kafka Pipeline Project (Golang + Kafka + Docker)

A production-style data pipeline built using **Golang**, **Apache Kafka**, and **Docker Compose**.

This project generates large datasets, pushes them to Kafka, consumes the data, sorts records based on multiple keys, and republishes the sorted output into separate Kafka topics.

---

# Tech Stack

- **Golang** – Core application logic
- **Apache Kafka** – Message streaming platform
- **Zookeeper** – Kafka dependency (single-node setup)
- **Docker Compose** – Container orchestration
- **Kafka-Go** (`segmentio/kafka-go`) – Kafka client library for Go

---

# Project Architecture

```text
Producer (Go Workers)
        ↓
   Kafka Topic: source
        ↓
Sorter Consumer (Go)
        ↓
 ┌───────────────┬───────────────┬────────────────┐
 ↓               ↓               ↓
Topic: id     Topic: name   Topic: continent
(sorted)      (sorted)      (sorted)


kafka_project/
│── cmd/
│   ├── producer/        # Kafka producer (worker mode)
│   └── sorter/          # Kafka consumer + sorting pipeline
│
│── internal/
│   ├── kafkautil/       # Kafka readers/writers
│   ├── sorter/          # Sorting logic
│   └── model/           # Record struct
│
│── docker-compose.yml
│── Dockerfile
│── go.mod
│── README.md


# How to Run
docker compose down -v
docker compose up -d --build
docker exec -it my_project-app-1 ./producer
docker exec -it my_project-app-1 ./sorter