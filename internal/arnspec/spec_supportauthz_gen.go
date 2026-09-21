// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: supportauthz
// Source: https://servicereference.us-east-1.amazonaws.com/v1/supportauthz/supportauthz.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "supportauthz_supportpermit", Service: "supportauthz", Resource: "supportpermit", Template: "arn:${Partition}:supportauthz:${Region}:${Account}:supportpermit/${ResourceId}"},
		{Name: "supportauthz_supportpermitrequest", Service: "supportauthz", Resource: "supportpermitrequest", Template: "arn:${Partition}:supportauthz:${Region}:${Account}:supportpermitrequest/${ResourceId}"},
	})
}
