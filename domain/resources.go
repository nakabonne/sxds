package domain

import (
	api "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/gogo/protobuf/proto"
)

// Resources are resources for xDS response. See below.
// https://www.envoyproxy.io/docs/envoy/latest/api-v2/api
type Resources struct {
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`

	Clusters  []*api.Cluster               `protobuf:"bytes,2,opt,name=clusters,proto3" json:"clusters,omitempty"`
	Endpoints []*api.ClusterLoadAssignment `protobuf:"bytes,3,opt,name=endpoints,proto3" json:"endpoints,omitempty"`
	Routes    []*api.RouteConfiguration    `protobuf:"bytes,4,opt,name=routes,proto3" json:"routes,omitempty"`
	Listeners []*api.Listener              `protobuf:"bytes,5,opt,name=listeners,proto3" json:"listeners,omitempty"`
}

func (m *Resources) Reset()         { *m = Resources{} }
func (m *Resources) String() string { return proto.CompactTextString(m) }
func (*Resources) ProtoMessage()    {}

var _ proto.Message = (*Resources)(nil)
