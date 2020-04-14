FROM golang:latest
RUN mkdir -p /app
WORKDIR /app
COPY main.go /app 
ENTRYPOINT [ "go", "run", "main.go" ]