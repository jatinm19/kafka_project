FROM golang:1.25

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o producer ./cmd/producer
RUN go build -o sorter ./cmd/sorter

CMD ["bash"]