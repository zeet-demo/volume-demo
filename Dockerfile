FROM golang:1.14

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build

EXPOSE 8080
VOLUME /data

ENV GIN_MODE=release

CMD ./volume-demo
