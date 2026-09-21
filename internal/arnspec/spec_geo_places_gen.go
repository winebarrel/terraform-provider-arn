// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: geo-places
// Source: https://servicereference.us-east-1.amazonaws.com/v1/geo-places/geo-places.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "geo_places_provider", Service: "geo-places", Resource: "provider", Template: "arn:${Partition}:geo-places:${Region}::provider/default"},
	})
}
