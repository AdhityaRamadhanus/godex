FROM golang:1.11

WORKDIR /go/src/github.com/AdhityaRamadhanus/godex
COPY . .

RUN make build

ENTRYPOINT ["./godex"]