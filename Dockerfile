FROM golang as builder

RUN mkdir /build

ADD . /build

WORKDIR /build
RUN GOOS=linux go build -o mrrobot ./cmd/main.go

FROM alpine

COPY --from=builder /build/mrrobot /mrrobot
ENTRYPOINT ["/mrrobot"]
