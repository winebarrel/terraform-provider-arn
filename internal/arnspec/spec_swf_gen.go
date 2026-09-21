// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: swf
// Source: https://servicereference.us-east-1.amazonaws.com/v1/swf/swf.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "swf_domain", Service: "swf", Resource: "domain", Template: "arn:${Partition}:swf::${Account}:/domain/${DomainName}"},
	})
}
