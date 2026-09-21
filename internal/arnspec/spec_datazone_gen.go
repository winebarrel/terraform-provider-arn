// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: datazone
// Source: https://servicereference.us-east-1.amazonaws.com/v1/datazone/datazone.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "datazone_domain", Service: "datazone", Resource: "domain", Template: "arn:${Partition}:datazone:${Region}:${Account}:domain/${DomainId}"},
	})
}
