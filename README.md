![build-master](https://github.com/clD11/go-whispers/workflows/build-master/badge.svg?branch=master)

# Go Whispers
Performs sentiment analysis on various sources and publishes cryptocurrency price.

* **Go Whispers Monorepo**
    * Getting started
    * Running locally
    * [Testing](#testing)

### Infrastructure
* AWS - EKS, SNS, SQS
* API gateway - [KrakenD](https://www.krakend.io/)

### Backend
* Golang

### UI
* NodeJs

### Testing
* Localstack for spinning up AWS and provisioned using Terraform
* Integration and e2e tests use Ginko