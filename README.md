# sxds

[![Release](https://img.shields.io/github/release/nakabonne/sxds.svg?style=flat-square)](https://github.com/nakabonne/sxds/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](/LICENSE.md)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/nakabonne/sxds)

Simple-xds for data-plane in ServiceMesh.  
This works as control-plane with the minimum function.

  
This README assumes you're familiar with the [data-plane-api](https://www.envoyproxy.io/docs/envoy/latest/configuration/overview/v2_overview) already

## Motivation
In realizing Service Mesh, istio is really a wonderful choice.  
But istio is big and complicated!  
If you want a control-plane with minimal functionality, sxds will be sufficiently effective.  
This xds supports any service proxy that conforms to data-plane-api.  
Currently only envoy is compliant, but envoy author Matt Klein says in [the blog](https://blog.envoyproxy.io/the-universal-data-plane-api-d15cec7a) that data-plane-api should be generic.

## Feature

- Return resoponse for DiscoveryRequest from data-plane.
- Cache DiscoveryResponse in memory

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
$ /path/to/sxds
```

### Required settings

#### Set resources

sxds caches resources and returns DiscoverResponse for data-plane.  
so you need to put the json file to the [cacher server](#cacher).  
   
Please create json file with reference to [sample](https://github.com/nakabonne/sxds/tree/master/sample/resource) and put json to the path `resources/{node_type}`.  
Put each [node type](#terms). 

```
$ curl -XPUT http://{IP_ADDRESS}:8082/resources/sidecar -d @sidecar.json
```

 
The json format has the same format as DiscoveryResponse in data-plane-api.  
For DiscoveryResponse, see [proto file](https://github.com/envoyproxy/data-plane-api/tree/master/envoy/api/v2) of data-plane-api.  
Also do not forget to update the version when updating.

#### Data-Plane settings


In order to suppress memory consumption, sxds cache resources for each node type.  
And sxds gets node type from node id, so you need to follow the naming convention.  
Please add node_type to prefix like "sidecar-app1" for naming node id of data-plane.  

##### e.g.) envoy  

[envoy-config]

Please set dynamic_resources and static_resources like [sample](https://github.com/nakabonne/sxds/blob/master/sample/envoy/envoy.yml).  

[specification of node id]  

Add node_type to the prefix of the name given to the `--service-node` option.  

```
$ envoy --service-node sidecar-1
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

- [ ] Make sxdsctl that is cli tool for put resources
- [ ] Make detailed documentation on resource json
- [ ] Automatic generation of resources json
- [ ] More test...
