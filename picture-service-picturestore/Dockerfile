#FROM node:7
#WORKDIR /app
#COPY package.json /app
#RUN npm install
#COPY . /app
#CMD node index.js
#EXPOSE 4000

# pulling a lightweight version of golang
FROM golang:1.8-alpine
MAINTAINER Mauricio Ribeiro <maumribeiro@gmail.com>
RUN apk --update add --no-cache git

# Copy the local package files to the container's workspace.
ADD . /go/src/picture-service-picturestore
WORKDIR /go/src/picture-service-picturestore

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get github.com/gorilla/handlers
RUN go get github.com/gorilla/mux

# Run the command by default when the container starts.
ENTRYPOINT ["/go/bin/picture-service-picturestore"]

# Document that the service listens on port 5000.
EXPOSE 5000