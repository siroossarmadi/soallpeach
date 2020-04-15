FROM golang:latest
RUN mkdir /app
WORKDIR /app
COPY main.go .
RUN go build main.go
ENTRYPOINT ["./main"]