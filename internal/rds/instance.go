package rds

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/pete911/aws-ip/internal"
	"time"
)

type Instance struct {
	AvailabilityZone       string
	CustomerOwnedIpEnabled bool
	DBClusterIdentifier    string
	DBInstanceArn          string
	DBInstanceClass        string
	DBInstanceIdentifier   string
	DBInstanceStatus       string
	DBName                 string
	VpcId                  string
	Subnets                []string
	DbInstancePort         int
	Endpoint               Endpoint
	Engine                 string
	EngineVersion          string
	InstanceCreateTime     time.Time
	Iops                   int
	MultiAZ                bool
	MultiTenant            bool
	NetworkType            string
	PubliclyAccessible     bool
	Tags                   map[string]string
}

func (i Instance) LookupIp() ([]string, error) {
	return internal.LookupIp(i.Endpoint.Address)
}

func toInstances(in []types.DBInstance) []Instance {
	var out []Instance
	for _, v := range in {
		out = append(out, toInstance(v))
	}
	return out
}

func toInstance(in types.DBInstance) Instance {
	var vpcId string
	var subnets []string
	if in.DBSubnetGroup != nil {
		vpcId = aws.ToString(in.DBSubnetGroup.VpcId)
		for _, subnet := range in.DBSubnetGroup.Subnets {
			subnets = append(subnets, aws.ToString(subnet.SubnetIdentifier))
		}
	}

	return Instance{
		AvailabilityZone:       aws.ToString(in.AvailabilityZone),
		CustomerOwnedIpEnabled: aws.ToBool(in.CustomerOwnedIpEnabled),
		DBClusterIdentifier:    aws.ToString(in.DBClusterIdentifier),
		DBInstanceArn:          aws.ToString(in.DBInstanceArn),
		DBInstanceClass:        aws.ToString(in.DBInstanceClass),
		DBInstanceIdentifier:   aws.ToString(in.DBInstanceIdentifier),
		DBInstanceStatus:       aws.ToString(in.DBInstanceStatus),
		DBName:                 aws.ToString(in.DBName),
		VpcId:                  vpcId,
		Subnets:                subnets,
		DbInstancePort:         int(aws.ToInt32(in.DbInstancePort)),
		Endpoint:               toEndpoint(in.Endpoint),
		Engine:                 aws.ToString(in.Engine),
		EngineVersion:          aws.ToString(in.EngineVersion),
		InstanceCreateTime:     aws.ToTime(in.InstanceCreateTime),
		Iops:                   int(aws.ToInt32(in.Iops)),
		MultiAZ:                aws.ToBool(in.MultiAZ),
		MultiTenant:            aws.ToBool(in.MultiTenant),
		NetworkType:            aws.ToString(in.NetworkType),
		PubliclyAccessible:     aws.ToBool(in.PubliclyAccessible),
		Tags:                   toTags(in.TagList),
	}
}

type Endpoint struct {
	Address      string
	HostedZoneId string
	Port         int
}

func toEndpoint(in *types.Endpoint) Endpoint {
	if in == nil {
		return Endpoint{}
	}

	return Endpoint{
		Address:      aws.ToString(in.Address),
		HostedZoneId: aws.ToString(in.HostedZoneId),
		Port:         int(aws.ToInt32(in.Port)),
	}
}

func toTags(in []types.Tag) map[string]string {
	out := make(map[string]string)
	for _, v := range in {
		out[aws.ToString(v.Key)] = aws.ToString(v.Value)
	}
	return out
}
