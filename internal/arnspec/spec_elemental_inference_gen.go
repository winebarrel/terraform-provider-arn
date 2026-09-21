// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: elemental-inference
// Source: https://servicereference.us-east-1.amazonaws.com/v1/elemental-inference/elemental-inference.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "elemental_inference_dictionary", Service: "elemental-inference", Resource: "dictionary", Template: "arn:${Partition}:elemental-inference:${Region}:${Account}:dictionary/${Id}"},
		{Name: "elemental_inference_feed", Service: "elemental-inference", Resource: "feed", Template: "arn:${Partition}:elemental-inference:${Region}:${Account}:feed/${Id}"},
	})
}
