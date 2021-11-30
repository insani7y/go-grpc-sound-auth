# The base go-image
FROM golang:latest

RUN apt-get update
RUN apt-get upgrade -y && apt-get -y install curl build-essential supervisor sudo

RUN curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
RUN apt-get update
RUN apt-get install -y migrate

# Create a directory for the app
RUN mkdir /app

# Copy all files from the current directory to the app directory
COPY . /app

# Set working directory
WORKDIR /app

RUN go mod download

RUN make
RUN make grpc

RUN chmod +x /app/entrypoint.sh

ENTRYPOINT [ "./entrypoint.sh" ]