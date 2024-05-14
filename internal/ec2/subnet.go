package ec2

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type Subnets []Subnet

func (s Subnets) BySubnetIds(ids []string) []Subnet {
	if len(ids) == 0 {
		return nil
	}

	idsSet := make(map[string]struct{})
	for _, id := range ids {
		idsSet[id] = struct{}{}
	}

	var out []Subnet
	for _, v := range s {
		if _, ok := idsSet[v.SubnetId]; ok {
			out = append(out, v)
		}
	}
	return out
}

type Subnet struct {
	VpcId                   string
	OwnerId                 string
	SubnetId                string
	SubnetArn               string
	AvailabilityZone        string
	AvailabilityZoneId      string
	AvailableIpAddressCount int
	CidrBlock               string
	DefaultForAz            bool
	MapPublicIpOnLaunch     bool
	Tags                    map[string]string
}

func toSubnets(in []types.Subnet) Subnets {
	var out []Subnet
	for _, v := range in {
		out = append(out, toSubnet(v))
	}
	return out
}

func toSubnet(in types.Subnet) Subnet {
	return Subnet{
		VpcId:                   aws.ToString(in.VpcId),
		OwnerId:                 aws.ToString(in.OwnerId),
		SubnetId:                aws.ToString(in.SubnetId),
		SubnetArn:               aws.ToString(in.SubnetArn),
		AvailabilityZone:        aws.ToString(in.AvailabilityZone),
		AvailabilityZoneId:      aws.ToString(in.AvailabilityZoneId),
		AvailableIpAddressCount: int(aws.ToInt32(in.AvailableIpAddressCount)),
		CidrBlock:               aws.ToString(in.CidrBlock),
		DefaultForAz:            aws.ToBool(in.DefaultForAz),
		MapPublicIpOnLaunch:     aws.ToBool(in.MapPublicIpOnLaunch),
		Tags:                    toTags(in.Tags),
	}
}
