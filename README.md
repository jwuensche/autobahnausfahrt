# autobahnausfahrt

![https://goreportcard.com/report/github.com/jwuensche/autobahnausfahrt](https://goreportcard.com/badge/github.com/jwuensche/autobahnausfahrt)

API exporter for the `autobahnkreuz` metrics API

`autobahnausfahrt` is designed to open any scrapable interface for monitoring databases (currently only prometheus)

## Configuration

Following environmeny variables are support to configure the exporter:

| flag   |  use | applicable values|
| ------|------|------------------|
| `INTERCONNECT_PORT` | define the port the router mertrics api is exposed | `7070` |
| `INTERCONNECT_ADDRESS` | define the current address the router is reachable on | `localhost` |
| `AUSFAHRT_PORT` | which port the scraper interface is exposed | `9101` |
| `AUSFAHRT_ROUTE` | on which route the interface is addressable | `/metrics` |
| `AUSFAHRT_CERT` | TLS certificate, if this value is set TLS will be used | `ausfahrt.cert` |
| `AUSFAHRT_KEY` | TLS key | `ausfahrt.key` |

## Defining own Interfaces

`autobahnausfahrt` is designed to use struct tags to allow a simple creation of new fields or structures. For Prometheus this tag is `prom`. The prometheus export is already provided, but others like InfluxDB s.o. have to be added.

## Building
Bulding from the provided `Dockerfile` will build a container based on the current `upstream/master`. Any local building can either be done by simply creating a binary, or by using the `DockerfileLocal` to create a `test` tagged container.

```
$ docker build -t autobahnausfahrt:test -f DockerfileLocal .
```
