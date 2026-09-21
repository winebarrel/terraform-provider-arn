// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: vpce
// Source: https://servicereference.us-east-1.amazonaws.com/v1/vpce/vpce.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "vpce_vpc_endpoint", Service: "vpce", Resource: "vpc-endpoint", Template: "arn:${Partition}:ec2:${Region}:${Account}:vpc-endpoint/${VpcEndpointId}"},
		{Name: "vpce_vpc_endpoint_service", Service: "vpce", Resource: "vpc-endpoint-service", Template: "arn:${Partition}:ec2:${Region}:${Account}:vpc-endpoint-service/${VpcEndpointServiceId}"},
	})
}
