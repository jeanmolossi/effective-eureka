FROM golang:1.18.4-alpine3.16

WORKDIR /usr/go

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install github.com/cespare/reflex@latest
RUN go install golang.org/x/tools/cmd/godoc@latest

RUN go mod tidy

RUN go build -o catalog

EXPOSE 8080

CMD [ "./catalog" ]
