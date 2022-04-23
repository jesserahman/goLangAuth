FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./app
COPY go.sum ./app
RUN go mod download

COPY *.go ./app

RUN go build -o /goLangAuth

EXPOSE 8081

CMD [ "/goLangAuth" ]