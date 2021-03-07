![build-master](https://github.com/clD11/go-whispers/workflows/build-master/badge.svg?branch=master)

# Whispers Roadmap PoC
- [x] Decide between mono/polyrepo and a basic project layout🤺
- [X] Implement basic Golang service and containerize 🐳
- [x] [Select API Gateway](#api-gateway) 🐙
- [x] Setup and containerize UI (react) 📺
- [x] [Select messaging technology](#messaging) 📨
- [ ] Service discovery (is this needed)
- [ ] Sentiment Analysis thinking AWS Comprehend or resurrecting old python project

## API Gateway <a name="api-gateway"/>
Many choices out there but have gone with [KrakenD](https://www.krakend.io/)

Besides sharing the same name as my favorite crypto exchange and a fairly impactful rum KrakenD is built using Go, is highly performant and its sweet spot seems to be backend service aggregation which is what I am looking for, this combined with near zero complexity and a declarative configuration makes it an easy choice.

## AWS SNS and SQS <a name="messaging"/>
* Decided on SNS and SQS over Kafka as its simpler to get up and running 
* For testing will use dockerised [elasticmq](https://github.com/softwaremill/elasticmq) to mock out SNS (trying to avoid overhead of [localstack](https://github.com/localstack/localstack) and keep things fast)

_Whisper initial ideas_ 
![whispers-design](whispers-design.jpg)
