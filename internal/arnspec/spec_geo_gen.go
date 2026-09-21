// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: geo
// Source: https://servicereference.us-east-1.amazonaws.com/v1/geo/geo.json
// Functions: 7
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "geo_api_key", Service: "geo", Resource: "api-key", Template: "arn:${Partition}:geo:${Region}:${Account}:api-key/${KeyName}"},
		{Name: "geo_geofence_collection", Service: "geo", Resource: "geofence-collection", Template: "arn:${Partition}:geo:${Region}:${Account}:geofence-collection/${GeofenceCollectionName}"},
		{Name: "geo_job", Service: "geo", Resource: "job", Template: "arn:${Partition}:geo:${Region}:${Account}:job/${JobId}"},
		{Name: "geo_map", Service: "geo", Resource: "map", Template: "arn:${Partition}:geo:${Region}:${Account}:map/${MapName}"},
		{Name: "geo_place_index", Service: "geo", Resource: "place-index", Template: "arn:${Partition}:geo:${Region}:${Account}:place-index/${IndexName}"},
		{Name: "geo_route_calculator", Service: "geo", Resource: "route-calculator", Template: "arn:${Partition}:geo:${Region}:${Account}:route-calculator/${CalculatorName}"},
		{Name: "geo_tracker", Service: "geo", Resource: "tracker", Template: "arn:${Partition}:geo:${Region}:${Account}:tracker/${TrackerName}"},
	})
}
