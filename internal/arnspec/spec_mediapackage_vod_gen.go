// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: mediapackage-vod
// Source: https://servicereference.us-east-1.amazonaws.com/v1/mediapackage-vod/mediapackage-vod.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "mediapackage_vod_assets", Service: "mediapackage-vod", Resource: "assets", Template: "arn:${Partition}:mediapackage-vod:${Region}:${Account}:assets/${AssetIdentifier}"},
		{Name: "mediapackage_vod_packaging_configurations", Service: "mediapackage-vod", Resource: "packaging-configurations", Template: "arn:${Partition}:mediapackage-vod:${Region}:${Account}:packaging-configurations/${PackagingConfigurationIdentifier}"},
		{Name: "mediapackage_vod_packaging_groups", Service: "mediapackage-vod", Resource: "packaging-groups", Template: "arn:${Partition}:mediapackage-vod:${Region}:${Account}:packaging-groups/${PackagingGroupIdentifier}"},
	})
}
