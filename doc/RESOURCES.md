# Resources document

You need to send json for [DiscoveryResponse](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/discovery.proto#discoveryresponse) to cacher.  
Please make a json file for each [node type](https://github.com/nakabonne/sxds#terms).  
The following fields must be set.  

- Four fields required for [DiscoveryResponse](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/discovery.proto#discoveryresponse)
- Version field

## Four fields required for DiscoveryResponse

- [listenrs](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/lds.proto#envoy-api-msg-listener)
- [clusters](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/cds.proto#cluster)
- [endpoints](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/eds.proto#envoy-api-msg-clusterloadassignment)
- [routes](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/rds.proto#routeconfiguration)

## Version field

You need to update the version each time you update resources.  
Anything is acceptable as long as it's a string. (e.g. "v1.1")

## Sample

- [sidecar](https://github.com/nakabonne/sxds/blob/master/sample/resource/sidecar.json)
