## base image we need for our go application

FROM golang:1.19.0-alpine3.16

## create an /app directory within our image that will hold our application source files

RUN mkdir /app

## copy everything in the root directory into our /app directory

ADD . /app

## specify that we now wish to execute any further commands inside our /app directory

WORKDIR /app

## go mod download command to pull in any dependencies

RUN go clean --modcache
RUN go mod download

## we run go build to compile the binary executable of our Go program

RUN go build -o main .

## start command which kicks off newly created binary executable

CMD ["/app/main"]