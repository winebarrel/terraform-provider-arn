// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: aiops
// Source: https://servicereference.us-east-1.amazonaws.com/v1/aiops/aiops.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "aiops_investigation_group", Service: "aiops", Resource: "investigation-group", Template: "arn:${Partition}:aiops:${Region}:${Account}:investigation-group/${InvestigationGroupId}"},
	})
}
