// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: healthlake
// Source: https://servicereference.us-east-1.amazonaws.com/v1/healthlake/healthlake.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "healthlake_data_transformation_profile", Service: "healthlake", Resource: "dataTransformationProfile", Template: "arn:${Partition}:healthlake:${Region}:${Account}:dataTransformationProfile/${ProfileId}"},
		{Name: "healthlake_datastore", Service: "healthlake", Resource: "datastore", Template: "arn:${Partition}:healthlake:${Region}:${Account}:datastore/fhir/${DatastoreId}"},
	})
}
