// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: securitylake
// Source: https://servicereference.us-east-1.amazonaws.com/v1/securitylake/securitylake.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "securitylake_data_lake", Service: "securitylake", Resource: "data-lake", Template: "arn:${Partition}:securitylake:${Region}:${Account}:data-lake/default"},
		{Name: "securitylake_subscriber", Service: "securitylake", Resource: "subscriber", Template: "arn:${Partition}:securitylake:${Region}:${Account}:subscriber/${SubscriberId}"},
	})
}
