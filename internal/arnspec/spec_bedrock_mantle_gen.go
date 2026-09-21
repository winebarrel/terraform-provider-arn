// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: bedrock-mantle
// Source: https://servicereference.us-east-1.amazonaws.com/v1/bedrock-mantle/bedrock-mantle.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "bedrock_mantle_customized_model", Service: "bedrock-mantle", Resource: "customized-model", Template: "arn:${Partition}:bedrock-mantle:${Region}:${Account}:customized-model/${ResourceId}"},
		{Name: "bedrock_mantle_project", Service: "bedrock-mantle", Resource: "project", Template: "arn:${Partition}:bedrock-mantle:${Region}:${Account}:project/${ResourceId}"},
		{Name: "bedrock_mantle_reservation", Service: "bedrock-mantle", Resource: "reservation", Template: "arn:${Partition}:bedrock-mantle:${Region}:${Account}:reservation/${ResourceId}"},
	})
}
