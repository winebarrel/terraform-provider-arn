// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: route53
// Source: https://servicereference.us-east-1.amazonaws.com/v1/route53/route53.json
// Functions: 9
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "route53_change", Service: "route53", Resource: "change", Template: "arn:${Partition}:route53:::change/${Id}"},
		{Name: "route53_cidrcollection", Service: "route53", Resource: "cidrcollection", Template: "arn:${Partition}:route53:::cidrcollection/${Id}"},
		{Name: "route53_delegationset", Service: "route53", Resource: "delegationset", Template: "arn:${Partition}:route53:::delegationset/${Id}"},
		{Name: "route53_healthcheck", Service: "route53", Resource: "healthcheck", Template: "arn:${Partition}:route53:::healthcheck/${Id}"},
		{Name: "route53_hostedzone", Service: "route53", Resource: "hostedzone", Template: "arn:${Partition}:route53:::hostedzone/${Id}"},
		{Name: "route53_queryloggingconfig", Service: "route53", Resource: "queryloggingconfig", Template: "arn:${Partition}:route53:::queryloggingconfig/${Id}"},
		{Name: "route53_trafficpolicy", Service: "route53", Resource: "trafficpolicy", Template: "arn:${Partition}:route53:::trafficpolicy/${Id}"},
		{Name: "route53_trafficpolicyinstance", Service: "route53", Resource: "trafficpolicyinstance", Template: "arn:${Partition}:route53:::trafficpolicyinstance/${Id}"},
		{Name: "route53_vpc", Service: "route53", Resource: "vpc", Template: "arn:${Partition}:ec2:${Region}:${Account}:vpc/${VpcId}"},
	})
}
