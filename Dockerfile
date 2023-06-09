FROM golang:1.20.3-alpine3.17

WORKDIR /jlpt-notify

COPY ./ ./

RUN go mod download

RUN go build -o /tmp/jlpt-notify

ENV TERM=xterm256color

CMD [ "/tmp/jlpt-notify" ]
