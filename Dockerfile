# syntax=docker/dockerfile:1
FROM golang:1.20

ENV APP_MAIN_ENTRY_PATH=cmd/logistics
ENV APP_BINARY_NAME=coop_logistics
ENV SERVICE_PORT 50051

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -a -v -o "${APP_MAIN_ENTRY_PATH}/${APP_BINARY_NAME}" "./${APP_MAIN_ENTRY_PATH}"

EXPOSE $SERVICE_PORT
CMD ["sh", "-c", "./${APP_MAIN_ENTRY_PATH}/${APP_BINARY_NAME}"]

