![build-master](https://github.com/clD11/go-whispers/workflows/build-master/badge.svg?branch=master)

# Whispers Roadmap PoC

- [x] Decide between mono/polyrepo and a basic project layout🤺
- [X] Implement basic Golang service and containerize 🐳
- [x] [Select API Gateway](#api-gateway) 🐙
- [x] Setup and containerize UI (react)
- [ ] Select messaging technology (thinking Kafka)
- [ ] Service discovery (is this needed)
- [ ] Sentiment Analysis thinking AWS Comprehend or resurrecting old python project

## API Gateway <a name="api-gateway"></a>
Many choices out there but have gone with [KrakenD](https://www.krakend.io/)

Besides sharing the same name as my favorite crypto exchange and a fairly impactful rum KrakenD is built using Go, is highly performant and its sweet spot seems to be backend service aggregation which is what I am looking for, this combined with near zero complexity and a declarative configuration makes it an easy choice.

Summary -

* On the-fly API creation using upstream services, with cross-cutting concerns (api gateway).
* Not a proxy, although it can be used as one.
* No node coordination
* No synchronization needed
* Zero complexity (docker container with a configuration file)
* No challenges for Multi-region
* Declarative configuration
* Immutable infrastructure
* Runs on micro and small machines in production without issues.
* Customizations in Go, Lua, CEL, and Martian DSL

_Whisper initial ideas_ 
![whispers-design](whispers-design.jpg)