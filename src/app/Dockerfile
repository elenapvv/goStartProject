FROM golang:1.19

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN apt-get -y update && apt-get -y upgrade && apt-get -y install openssl libssl-dev
RUN go build -v -o /usr/local/bin ./...

CMD ["app"]
