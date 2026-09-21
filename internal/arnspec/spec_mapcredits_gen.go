// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: mapcredits
// Source: https://servicereference.us-east-1.amazonaws.com/v1/mapcredits/mapcredits.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "mapcredits_agreement", Service: "mapcredits", Resource: "agreement", Template: "arn:${Partition}:mapcredits:::${Agreement}/${AgreementId}"},
	})
}
