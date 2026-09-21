// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: servicediscovery
// Source: https://servicereference.us-east-1.amazonaws.com/v1/servicediscovery/servicediscovery.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "servicediscovery_namespace", Service: "servicediscovery", Resource: "namespace", Template: "arn:${Partition}:servicediscovery:${Region}:${Account}:namespace/${NamespaceId}"},
		{Name: "servicediscovery_service", Service: "servicediscovery", Resource: "service", Template: "arn:${Partition}:servicediscovery:${Region}:${Account}:service/${ServiceId}"},
	})
}
