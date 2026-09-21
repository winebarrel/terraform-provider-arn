// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: medical-imaging
// Source: https://servicereference.us-east-1.amazonaws.com/v1/medical-imaging/medical-imaging.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "medical_imaging_datastore", Service: "medical-imaging", Resource: "datastore", Template: "arn:${Partition}:medical-imaging:${Region}:${Account}:datastore/${DatastoreId}"},
		{Name: "medical_imaging_imageset", Service: "medical-imaging", Resource: "imageset", Template: "arn:${Partition}:medical-imaging:${Region}:${Account}:datastore/${DatastoreId}/imageset/${ImageSetId}"},
	})
}
