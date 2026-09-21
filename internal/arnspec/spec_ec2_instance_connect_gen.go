// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: ec2-instance-connect
// Source: https://servicereference.us-east-1.amazonaws.com/v1/ec2-instance-connect/ec2-instance-connect.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "ec2_instance_connect_instance", Service: "ec2-instance-connect", Resource: "instance", Template: "arn:${Partition}:ec2:${Region}:${Account}:instance/${InstanceId}"},
		{Name: "ec2_instance_connect_instance_connect_endpoint", Service: "ec2-instance-connect", Resource: "instance-connect-endpoint", Template: "arn:${Partition}:ec2:${Region}:${Account}:instance-connect-endpoint/${InstanceConnectEndpointId}"},
	})
}
