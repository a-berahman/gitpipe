FROM golang:1.15-alpine AS build
#Install git
RUN apk add --no-cache git
#Get the package from a GitHub repository
RUN go get github.com/a-berahman/gitpipe
WORKDIR /go/src/github.com/a-berahman/gitpipe
# Build the project and send the output to /bin/GitPipe 
RUN go build -o /bin/GitPipe

FROM golang:1.15-alpine
#Copy the build's output binary from the previous build container
COPY --from=build /bin/GitPipe /bin/GitPipe
ENTRYPOINT ["/bin/GitPipe"]