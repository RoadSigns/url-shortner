# syntax=docker/dockerfile:1

FROM golang:1.18-alpine3.15

ENV GO111MODULE=on

## We create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
ADD . /app
## We specify that we now wish to execute
## any further commands inside our /app
## directory
WORKDIR /app

# download Go modules and dependencies
RUN go mod download

# compile application
RUN go build -o /godocker

# tells Docker that the container listens on specified network ports at runtime
EXPOSE 8080

# command to be used to execute when the image is used to start a container
CMD [ "/godocker" ]