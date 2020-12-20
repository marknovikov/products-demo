# Products service demo

## Description

Demo gRPC server implementation behind NGINX load balancer.

### Service implements 2 methods:

- `Fetch(url)` loads external \*.csv listing of available products (name; price) by provided url, stores products in the DB, updating prices and meta as needed. Nodejs mock service is provided to mock the \*.csv file providing external server.
- `List(paging, sorting)` lists all products, possibly with keyset paging and sorting by allowed fields.

To inspect API check out `./api/products.proto` API definition and `./requests.example` API usage examples.

## Used stack

- Go 1.15
- Protobuf
- gRPC
- MongoDB
- Node.js
- NGINX
- Docker
- Docker-Compose
- Make
- grpcurl

## Prerequisites

1. Make
1. Docker
1. Docker-compose

And grpcurl in order to test out the API manually.

All the remaining dependencies are completely dockerised.

## Usage

1. Install [grpcurl](https://github.com/fullstorydev/grpcurl)
1. cd to project dir
1. Compile proto description for grpcurl: `make protoset`
1. Spin up containers: `make up`
1. Try grpcurl requests. Examples of working requests may be found in `./requests.example`
1. Shut down the containers `make down`
