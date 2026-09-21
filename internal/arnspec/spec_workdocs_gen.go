// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: workdocs
// Source: https://servicereference.us-east-1.amazonaws.com/v1/workdocs/workdocs.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "workdocs_organization", Service: "workdocs", Resource: "organization", Template: "arn:${Partition}:workdocs:${Region}:${Account}:organization/${ResourceId}"},
	})
}
