# sxds

[![CircleCI](https://circleci.com/gh/nakabonne/sxds.svg?style=svg)](https://circleci.com/gh/nakabonne/sxds)
[![Release](https://img.shields.io/github/release/nakabonne/sxds.svg?style=flat-square)](https://github.com/nakabonne/sxds/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](/LICENSE.md)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/nakabonne/sxds)

Simple xds for data-plane in ServiceMesh.  
Sxds enables service discovery, dynamic updates to load balancing pools and routing tables, and supports any data-plane that conforms to data-plane-api.  
The communication protocol with data-plane is supported only by gRPC.

  
This README assumes you're familiar with the [data-plane-api](https://www.envoyproxy.io/docs/envoy/latest/configuration/overview/v2_overview) already.

## Motivation
Using an orchestration tool makes it possible to achieve Service Mesh relatively easily.  
But if you don't use it, you can't benefit from a great tool such as Istio!  
If you want to realize Service Mesh in such an environment, sxds is one of effective methods.  

## Feature

- Provides policy and configuration for all of the running data planes
  - Listener discovery service ([LDS](https://www.envoyproxy.io/docs/envoy/latest/configuration/listeners/lds))
  - Cluster discovery service ([CDS](https://www.envoyproxy.io/docs/envoy/latest/configuration/cluster_manager/cds))
  - Route discovery service ([RDS](https://www.envoyproxy.io/docs/envoy/latest/configuration/http_conn_man/rds))
  - Endpoint discovery service ([EDS](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/eds.proto#envoy-api-file-envoy-api-v2-eds-proto))

## Installation

### Binaries

Please see [Github Releases](https://github.com/nakabonne/sxds/releases).

### From source

```
$ go get -u github.com/nakabonne/sxds
$ go install github.com/nakabonne/sxds/cmd/sxds
```

## Usage

### Run

Two servers, [xds](#xds) and [cacher](#cacher) will listen.

```sh
$ sxds
```

### Required settings

There are only two things you have to do beforehand.

#### Set resources

sxds caches resources and returns DiscoverResponse for data-plane.  
so you need to send the json file to the [cacher server](#cacher).  
Please create json while referring to the [document](https://github.com/nakabonne/sxds/tree/master/doc/RESOURCES.md).

```
$ curl -XPUT http://{IP_ADDRESS}:8082/resources/sidecar -d @sidecar.json
```

#### Data-Plane settings


In order to suppress memory consumption, sxds cache resources for each node type.  
And sxds gets node type from node id, so you need to follow the naming convention.  
Please add node_type to prefix like "sidecar-app1" for naming node id of data-plane.  

##### envoy  

[envoy-config]

Add sxds cluster to static_resources and specify it in dynamic_resources.  
See [sample](https://github.com/nakabonne/sxds/blob/master/sample/envoy/envoy.yml).  

[specification of node id]  

Add node_type to the prefix of the name given to the `--service-node` option.  

```
$ envoy --service-node sidecar-app1 --service-cluster app1
```

### Optional settings

```sh
SXDS_PRODUCTION=false # default: false
SXDS_ADS_MODE=false   # default: false
SXDS_XDS_PORT=8081    # default: 8081
SXDS_CACHER_PORT=8082 # default: 8082
```

For ADS mode, please click [here](https://github.com/envoyproxy/data-plane-api/blob/master/XDS_PROTOCOL.md#aggregated-discovery-services-ads)

## Architecture

![architecture](https://github.com/nakabonne/sxds/blob/master/media/architecture.png) 

### xDS
gRPC server that return response to data-plane.

### cacher

REST server that caches resources.


## Terms

| term | meaning |
|:----------|:-----------|
|node type|The role of node which put each data-plane(e.g. sidecar, router)|
|resources|Data to use for DiscoveryResponse(e.g. listeners, clusters) |

## TODO

- [x] Make detailed documentation on resource json
- [ ] Make sxdsctl that is cli tool for put resources
- [ ] Update to be available synchronous (long) polling via REST endpoints 
- [ ] Automatic generation of resources json
- [ ] More test...
