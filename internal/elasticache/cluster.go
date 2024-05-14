package elasticache

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/pete911/aws-ip/internal"
	"time"
)

type Cluster struct {
	ARN                       string
	AtRestEncryptionEnabled   bool
	AuthTokenEnabled          bool
	AuthTokenLastModifiedDate time.Time
	CacheClusterCreateTime    time.Time
	CacheClusterId            string
	CacheClusterStatus        string
	CacheNodeType             string
	CacheNodes                []Node
	CacheSubnetGroupName      string
	ConfigurationEndpoint     Endpoint
	Engine                    string
	EngineVersion             string
	IpDiscovery               string
	NetworkType               string
	NumCacheNodes             int
}

func toClusters(in []types.CacheCluster) []Cluster {
	var out []Cluster
	for _, v := range in {
		out = append(out, toCluster(v))
	}
	return out
}

func toCluster(in types.CacheCluster) Cluster {
	return Cluster{
		ARN:                       aws.ToString(in.ARN),
		AtRestEncryptionEnabled:   aws.ToBool(in.AtRestEncryptionEnabled),
		AuthTokenEnabled:          aws.ToBool(in.AuthTokenEnabled),
		AuthTokenLastModifiedDate: aws.ToTime(in.AuthTokenLastModifiedDate),
		CacheClusterCreateTime:    aws.ToTime(in.CacheClusterCreateTime),
		CacheClusterId:            aws.ToString(in.CacheClusterId),
		CacheClusterStatus:        aws.ToString(in.CacheClusterStatus),
		CacheNodeType:             aws.ToString(in.CacheNodeType),
		CacheNodes:                toNodes(in.CacheNodes),
		CacheSubnetGroupName:      aws.ToString(in.CacheSubnetGroupName),
		ConfigurationEndpoint:     toEndpoint(in.ConfigurationEndpoint),
		Engine:                    aws.ToString(in.Engine),
		EngineVersion:             aws.ToString(in.EngineVersion),
		IpDiscovery:               string(in.IpDiscovery),
		NetworkType:               string(in.NetworkType),
		NumCacheNodes:             int(aws.ToInt32(in.NumCacheNodes)),
	}
}

type Node struct {
	CacheNodeCreateTime time.Time
	CacheNodeId         string
	CacheNodeStatus     string
	Endpoint            Endpoint
	SourceCacheNodeId   string
}

func (n Node) LookupIp() ([]string, error) {
	return internal.LookupIp(n.Endpoint.Address)
}

func toNodes(in []types.CacheNode) []Node {
	var out []Node
	for _, v := range in {
		out = append(out, toNode(v))
	}
	return out
}

func toNode(in types.CacheNode) Node {
	return Node{
		CacheNodeCreateTime: aws.ToTime(in.CacheNodeCreateTime),
		CacheNodeId:         aws.ToString(in.CacheNodeId),
		CacheNodeStatus:     aws.ToString(in.CacheNodeStatus),
		Endpoint:            toEndpoint(in.Endpoint),
		SourceCacheNodeId:   aws.ToString(in.SourceCacheNodeId),
	}
}

type Endpoint struct {
	Address string
	Port    int
}

func toEndpoint(in *types.Endpoint) Endpoint {
	if in == nil {
		return Endpoint{}
	}
	return Endpoint{
		Address: aws.ToString(in.Address),
		Port:    int(aws.ToInt32(in.Port)),
	}
}
