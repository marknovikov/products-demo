FROM golang:1.15 AS builder
WORKDIR /products-demo
COPY . .
RUN apt -qq update 2> /dev/null && \
    apt -qq install unzip 2> /dev/null && \
    apt -qq install make 2> /dev/null && \
    apt -qq install curl 2> /dev/null && \
    apt -qq install git 2> /dev/null
RUN make build

FROM alpine:latest AS runner
COPY --from=builder /products-demo/bin/products /products
EXPOSE 8080
ENTRYPOINT [ "/products" ]
