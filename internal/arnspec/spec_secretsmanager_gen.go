// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: secretsmanager
// Source: https://servicereference.us-east-1.amazonaws.com/v1/secretsmanager/secretsmanager.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "secretsmanager_secret", Service: "secretsmanager", Resource: "Secret", Template: "arn:${Partition}:secretsmanager:${Region}:${Account}:secret:${SecretId}"},
	})
}
