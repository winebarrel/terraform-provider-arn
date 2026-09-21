// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: workmailmessageflow
// Source: https://servicereference.us-east-1.amazonaws.com/v1/workmailmessageflow/workmailmessageflow.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "workmailmessageflow_raw_message", Service: "workmailmessageflow", Resource: "RawMessage", Template: "arn:${Partition}:workmailmessageflow:${Region}:${Account}:message/${OrganizationId}/${Context}/${MessageId}"},
	})
}
