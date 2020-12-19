# Products service demo

## Requirements

1. Docker
2. Docker-compose

## Usage

1. Install [grpcurl](https://github.com/fullstorydev/grpcurl)
1. cd to project dir
1. Compile proto description for grpcurl: `make protoset`
1. Spin up containers: `make up`
1. Try grpcurl requests. Examples of working requests may be found in `./requests.example`
1. Shut down the containers `make down`
