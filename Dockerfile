FROM golang:1.15 AS build


RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go get github.com/a-berahman/gitpipe
WORKDIR /go/src/github.com/a-berahman/gitpipe
RUN go build -o /app/gitpipe
# RUN GOPATH=/usr/go CGO_ENABLED=0 go build -o gitpipe .

FROM alpine:3.12

COPY --from=build /app/gitpipe /app/entrypoint.sh /app/

RUN apk update && \
    apk add --update bash && \
    apk add --update tzdata && \
    chmod +x /app/gitpipe /app/entrypoint.sh


ENTRYPOINT ["./entrypoint.sh"]
CMD ["serve"]
