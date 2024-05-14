package ec2

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type Vpcs []Vpc

func (v Vpcs) ById(id string) Vpc {
	for _, vpc := range v {
		if vpc.VpcId == id {
			return vpc
		}
	}
	return Vpc{}
}

type Vpc struct {
	VpcId         string
	OwnerId       string
	CidrBlock     string
	DhcpOptionsId string
	IsDefault     bool
	Tags          map[string]string
}

func toVpcs(in []types.Vpc) Vpcs {
	var out []Vpc
	for _, v := range in {
		out = append(out, toVpc(v))
	}
	return out
}

func toVpc(in types.Vpc) Vpc {
	return Vpc{
		VpcId:         aws.ToString(in.VpcId),
		OwnerId:       aws.ToString(in.OwnerId),
		CidrBlock:     aws.ToString(in.CidrBlock),
		DhcpOptionsId: aws.ToString(in.DhcpOptionsId),
		IsDefault:     aws.ToBool(in.IsDefault),
		Tags:          toTags(in.Tags),
	}
}
