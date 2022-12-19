FROM golang:1.19.3-alpine

RUN go version
ENV GOPATH=/
ENV APP_PATH=/home/src

EXPOSE 8080
WORKDIR $APP_PATH

COPY ./ $APP_PATH

RUN go mod download
RUN go build -o facade-service $APP_PATH/cmd/main.go

ENTRYPOINT ["./facade-service"]