FROM golang:1.20.3-alpine3.17

WORKDIR /jlpt-notify

COPY ./go.mod ./

RUN go mod download

CMD [ "go", "run", "." ]
