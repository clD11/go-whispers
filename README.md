![build-master](https://github.com/clD11/go-whispers/workflows/build-master/badge.svg?branch=master)

# Go Whispers
Performs sentiment analysis on various sources and publishes cryptocurrency price.

### Infrastructure
* AWS - EKS, SNS, SQS
* API gateway - [KrakenD](https://www.krakend.io/)

### Backend
* Golang

### UI
* NodeJs

### Testing
* Localstack for spinning up AWS and infra applied using Terraform, this allows for testing locally
* Ginko for integration and e2e