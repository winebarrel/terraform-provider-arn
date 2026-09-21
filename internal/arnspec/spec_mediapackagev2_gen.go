// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: mediapackagev2
// Source: https://servicereference.us-east-1.amazonaws.com/v1/mediapackagev2/mediapackagev2.json
// Functions: 6
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "mediapackagev2_channel", Service: "mediapackagev2", Resource: "Channel", Template: "arn:${Partition}:mediapackagev2:${Region}:${Account}:channelGroup/${ChannelGroupName}/channel/${ChannelName}"},
		{Name: "mediapackagev2_channel_group", Service: "mediapackagev2", Resource: "ChannelGroup", Template: "arn:${Partition}:mediapackagev2:${Region}:${Account}:channelGroup/${ChannelGroupName}"},
		{Name: "mediapackagev2_channel_policy", Service: "mediapackagev2", Resource: "ChannelPolicy", Template: "arn:${Partition}:mediapackagev2:${Region}:${Account}:channelGroup/${ChannelGroupName}/channel/${ChannelName}"},
		{Name: "mediapackagev2_harvest_job", Service: "mediapackagev2", Resource: "HarvestJob", Template: "arn:${Partition}:mediapackagev2:${Region}:${Account}:channelGroup/${ChannelGroupName}/channel/${ChannelName}/originEndpoint/${OriginEndpointName}/harvestJob/${HarvestJobName}"},
		{Name: "mediapackagev2_origin_endpoint", Service: "mediapackagev2", Resource: "OriginEndpoint", Template: "arn:${Partition}:mediapackagev2:${Region}:${Account}:channelGroup/${ChannelGroupName}/channel/${ChannelName}/originEndpoint/${OriginEndpointName}"},
		{Name: "mediapackagev2_origin_endpoint_policy", Service: "mediapackagev2", Resource: "OriginEndpointPolicy", Template: "arn:${Partition}:mediapackagev2:${Region}:${Account}:channelGroup/${ChannelGroupName}/channel/${ChannelName}/originEndpoint/${OriginEndpointName}"},
	})
}
