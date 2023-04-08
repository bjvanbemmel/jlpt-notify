FROM golang:1.20.3-alpine3.17

WORKDIR /jlpt-notify

COPY ./go.mod ./

RUN go mod download

RUN go build -o ./bin/jlpt-notify .

ENV TERM=xterm256color

CMD [ "./bin/jlpt-notify" ]
