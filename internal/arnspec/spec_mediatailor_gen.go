// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: mediatailor
// Source: https://servicereference.us-east-1.amazonaws.com/v1/mediatailor/mediatailor.json
// Functions: 7
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "mediatailor_channel", Service: "mediatailor", Resource: "channel", Template: "arn:${Partition}:mediatailor:${Region}:${Account}:channel/${ChannelName}"},
		{Name: "mediatailor_live_source", Service: "mediatailor", Resource: "liveSource", Template: "arn:${Partition}:mediatailor:${Region}:${Account}:liveSource/${SourceLocationName}/${LiveSourceName}"},
		{Name: "mediatailor_playback_configuration", Service: "mediatailor", Resource: "playbackConfiguration", Template: "arn:${Partition}:mediatailor:${Region}:${Account}:playbackConfiguration/${ResourceId}"},
		{Name: "mediatailor_prefetch_schedule", Service: "mediatailor", Resource: "prefetchSchedule", Template: "arn:${Partition}:mediatailor:${Region}:${Account}:prefetchSchedule/${ResourceId}"},
		{Name: "mediatailor_program", Service: "mediatailor", Resource: "program", Template: "arn:${Partition}:mediatailor:${Region}:${Account}:program/${ChannelName}/${ProgramName}"},
		{Name: "mediatailor_source_location", Service: "mediatailor", Resource: "sourceLocation", Template: "arn:${Partition}:mediatailor:${Region}:${Account}:sourceLocation/${SourceLocationName}"},
		{Name: "mediatailor_vod_source", Service: "mediatailor", Resource: "vodSource", Template: "arn:${Partition}:mediatailor:${Region}:${Account}:vodSource/${SourceLocationName}/${VodSourceName}"},
	})
}
