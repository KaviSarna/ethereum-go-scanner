FROM golang:1.12

WORKDIR /go/src/matic

COPY ./main.go .
COPY ./vendor/ .

RUN go get -u github.com/ethereum/go-ethereum
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/lib/pq

CMD ["go", "run", "main.go"]
