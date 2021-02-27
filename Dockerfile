FROM golang:1.15 AS build


RUN mkdir /app
ADD . /app/
WORKDIR /app


RUN go get github.com/a-berahman/gitpipe
WORKDIR /go/src/github.com/a-berahman/gitpipe
RUN ls /go/src/github.com/a-berahman/gitpipe
RUN ls /go/
RUN pwd
RUN GOPATH=/go GOBIN=/go/bin  go build -o /bin/gitpipe
# RUN go build -o /app/gitpipe
# RUN GOPATH=/usr/go CGO_ENABLED=0 go build -o gitpipe .

FROM alpine:3.12

COPY --from=build /bin/gitpipe /bin/entrypoint.sh /bin/

RUN apk update && \
    apk add --update bash && \
    apk add --update tzdata && \
    chmod +x /bin/gitpipe /bin/entrypoint.sh


ENTRYPOINT ["./entrypoint.sh"]
CMD ["serve"]
