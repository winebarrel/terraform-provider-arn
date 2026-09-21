// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: geo-routes
// Source: https://servicereference.us-east-1.amazonaws.com/v1/geo-routes/geo-routes.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "geo_routes_provider", Service: "geo-routes", Resource: "provider", Template: "arn:${Partition}:geo-routes:${Region}::provider/default"},
	})
}
