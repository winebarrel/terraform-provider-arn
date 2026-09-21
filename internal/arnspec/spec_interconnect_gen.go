// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: interconnect
// Source: https://servicereference.us-east-1.amazonaws.com/v1/interconnect/interconnect.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "interconnect_connection", Service: "interconnect", Resource: "connection", Template: "arn:${Partition}:interconnect:${Region}:${Account}:connection/${Id}"},
		{Name: "interconnect_environment", Service: "interconnect", Resource: "environment", Template: "arn:${Partition}:interconnect:${Region}:${Account}:environment/${Id}"},
	})
}
