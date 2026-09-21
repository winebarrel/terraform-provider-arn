// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: cloudsearch
// Source: https://servicereference.us-east-1.amazonaws.com/v1/cloudsearch/cloudsearch.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "cloudsearch_domain", Service: "cloudsearch", Resource: "domain", Template: "arn:${Partition}:cloudsearch:${Region}:${Account}:domain/${DomainName}"},
	})
}
