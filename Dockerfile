FROM golang:1.12
COPY . /go/src/github.com/gersakbogdan/fasttrackquiz
WORKDIR /go/src/github.com/gersakbogdan/fasttrackquiz
RUN go install ./cmd/quizer
