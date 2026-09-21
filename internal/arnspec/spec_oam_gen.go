// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: oam
// Source: https://servicereference.us-east-1.amazonaws.com/v1/oam/oam.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "oam_link", Service: "oam", Resource: "Link", Template: "arn:${Partition}:oam:${Region}:${Account}:link/${ResourceId}"},
		{Name: "oam_sink", Service: "oam", Resource: "Sink", Template: "arn:${Partition}:oam:${Region}:${Account}:sink/${ResourceId}"},
	})
}
