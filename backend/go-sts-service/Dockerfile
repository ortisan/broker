FROM golang:alpine
COPY . /usr/local/go/pkg/ortisan-broker/go-sts-service/
WORKDIR /usr/local/go/pkg/ortisan-broker/go-sts-service/
RUN go get -d -v
CMD ["go", "run", "main.go"]