FROM golang:1.14.2 AS dev
WORKDIR /work/go-cli-sandbox
ENV GO111MODULE=on

FROM golang AS builder
COPY ./ ./
RUN make prepare build-linux

FROM alpine AS app
COPY --from=builder /work/go-cli-sandbox/build/go-cli-sandbox-linux-amd64 /usr/local/bin/go-cli-sandbox
CMD ["go-cli-sandbox"]
