FROM golang:1.15 AS builder
RUN apt -qq update 2> /dev/null && \
    apt -qq install unzip 2> /dev/null && \
    apt -qq install make 2> /dev/null && \
    apt -qq install curl 2> /dev/null && \
    apt -qq install git 2> /dev/null
WORKDIR /products-demo
COPY go.mod go.sum dev.env Makefile ./
RUN make deps
COPY api api
RUN make proto
COPY . .
RUN make build-in-docker

FROM alpine:latest AS runner
EXPOSE 8080
COPY --from=builder /products-demo/bin/products /products
ENTRYPOINT [ "/products" ]
