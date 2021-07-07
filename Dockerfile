FROM golang:1.16.5

WORKDIR /go/src/app
COPY . .

RUN rm -f go.mod
RUN go mod init
RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]