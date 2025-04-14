build:
    go build -o bin/rinha cmd/app/main.go

run: build
    ./bin/rinha

postgres:
    docker compose up -d
