// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: ivs
// Source: https://servicereference.us-east-1.amazonaws.com/v1/ivs/ivs.json
// Functions: 12
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "ivs_ad_configuration", Service: "ivs", Resource: "Ad-Configuration", Template: "arn:${Partition}:ivs:${Region}:${Account}:ad-configuration/${ResourceId}"},
		{Name: "ivs_channel", Service: "ivs", Resource: "Channel", Template: "arn:${Partition}:ivs:${Region}:${Account}:channel/${ResourceId}"},
		{Name: "ivs_composition", Service: "ivs", Resource: "Composition", Template: "arn:${Partition}:ivs:${Region}:${Account}:composition/${ResourceId}"},
		{Name: "ivs_encoder_configuration", Service: "ivs", Resource: "Encoder-Configuration", Template: "arn:${Partition}:ivs:${Region}:${Account}:encoder-configuration/${ResourceId}"},
		{Name: "ivs_ingest_configuration", Service: "ivs", Resource: "Ingest-Configuration", Template: "arn:${Partition}:ivs:${Region}:${Account}:ingest-configuration/${ResourceId}"},
		{Name: "ivs_playback_key_pair", Service: "ivs", Resource: "Playback-Key-Pair", Template: "arn:${Partition}:ivs:${Region}:${Account}:playback-key/${ResourceId}"},
		{Name: "ivs_playback_restriction_policy", Service: "ivs", Resource: "Playback-Restriction-Policy", Template: "arn:${Partition}:ivs:${Region}:${Account}:playback-restriction-policy/${ResourceId}"},
		{Name: "ivs_public_key", Service: "ivs", Resource: "Public-Key", Template: "arn:${Partition}:ivs:${Region}:${Account}:public-key/${ResourceId}"},
		{Name: "ivs_recording_configuration", Service: "ivs", Resource: "Recording-Configuration", Template: "arn:${Partition}:ivs:${Region}:${Account}:recording-configuration/${ResourceId}"},
		{Name: "ivs_stage", Service: "ivs", Resource: "Stage", Template: "arn:${Partition}:ivs:${Region}:${Account}:stage/${ResourceId}"},
		{Name: "ivs_storage_configuration", Service: "ivs", Resource: "Storage-Configuration", Template: "arn:${Partition}:ivs:${Region}:${Account}:storage-configuration/${ResourceId}"},
		{Name: "ivs_stream_key", Service: "ivs", Resource: "Stream-Key", Template: "arn:${Partition}:ivs:${Region}:${Account}:stream-key/${ResourceId}"},
	})
}
