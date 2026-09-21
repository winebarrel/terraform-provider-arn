// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: voiceid
// Source: https://servicereference.us-east-1.amazonaws.com/v1/voiceid/voiceid.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "voiceid_domain", Service: "voiceid", Resource: "domain", Template: "arn:${Partition}:voiceid:${Region}:${Account}:domain/${DomainId}"},
	})
}
