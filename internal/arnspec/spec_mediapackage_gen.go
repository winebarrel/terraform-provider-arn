// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: mediapackage
// Source: https://servicereference.us-east-1.amazonaws.com/v1/mediapackage/mediapackage.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "mediapackage_channels", Service: "mediapackage", Resource: "channels", Template: "arn:${Partition}:mediapackage:${Region}:${Account}:channels/${ChannelIdentifier}"},
		{Name: "mediapackage_harvest_jobs", Service: "mediapackage", Resource: "harvest_jobs", Template: "arn:${Partition}:mediapackage:${Region}:${Account}:harvest_jobs/${HarvestJobIdentifier}"},
		{Name: "mediapackage_origin_endpoints", Service: "mediapackage", Resource: "origin_endpoints", Template: "arn:${Partition}:mediapackage:${Region}:${Account}:origin_endpoints/${OriginEndpointIdentifier}"},
	})
}
