// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: geo-maps
// Source: https://servicereference.us-east-1.amazonaws.com/v1/geo-maps/geo-maps.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "geo_maps_provider", Service: "geo-maps", Resource: "provider", Template: "arn:${Partition}:geo-maps:${Region}::provider/default"},
	})
}
