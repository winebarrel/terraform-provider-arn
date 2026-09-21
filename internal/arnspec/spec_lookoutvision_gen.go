// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: lookoutvision
// Source: https://servicereference.us-east-1.amazonaws.com/v1/lookoutvision/lookoutvision.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "lookoutvision_model", Service: "lookoutvision", Resource: "model", Template: "arn:${Partition}:lookoutvision:${Region}:${Account}:model/${ProjectName}/${ModelVersion}"},
		{Name: "lookoutvision_project", Service: "lookoutvision", Resource: "project", Template: "arn:${Partition}:lookoutvision:${Region}:${Account}:project/${ProjectName}"},
	})
}
