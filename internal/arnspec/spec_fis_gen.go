// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: fis
// Source: https://servicereference.us-east-1.amazonaws.com/v1/fis/fis.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "fis_action", Service: "fis", Resource: "action", Template: "arn:${Partition}:fis:${Region}:${Account}:action/${Id}"},
		{Name: "fis_experiment", Service: "fis", Resource: "experiment", Template: "arn:${Partition}:fis:${Region}:${Account}:experiment/${Id}"},
		{Name: "fis_experiment_template", Service: "fis", Resource: "experiment-template", Template: "arn:${Partition}:fis:${Region}:${Account}:experiment-template/${Id}"},
		{Name: "fis_safety_lever", Service: "fis", Resource: "safety-lever", Template: "arn:${Partition}:fis:${Region}:${Account}:safety-lever/${Id}"},
	})
}
