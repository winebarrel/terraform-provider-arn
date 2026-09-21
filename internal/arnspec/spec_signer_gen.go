// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: signer
// Source: https://servicereference.us-east-1.amazonaws.com/v1/signer/signer.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "signer_signing_job", Service: "signer", Resource: "signing-job", Template: "arn:${Partition}:signer:${Region}:${Account}:/signing-jobs/${JobId}"},
		{Name: "signer_signing_profile", Service: "signer", Resource: "signing-profile", Template: "arn:${Partition}:signer:${Region}:${Account}:/signing-profiles/${ProfileName}"},
	})
}
