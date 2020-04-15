FROM golang:latest
RUN mkdir /app
WORKDIR /app
COPY main.go /app
RUN go build main.go
ENTRYPOINT ["./main"]