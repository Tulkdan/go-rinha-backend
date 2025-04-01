build:
    go build -o bin/rinha

run: build
    ./bin/rinha

postgres:
    docker compose up -d
