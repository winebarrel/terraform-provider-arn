// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: chime
// Source: https://servicereference.us-east-1.amazonaws.com/v1/chime/chime.json
// Functions: 13
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "chime_app_instance", Service: "chime", Resource: "app-instance", Template: "arn:${Partition}:chime:${Region}:${AccountId}:app-instance/${AppInstanceId}"},
		{Name: "chime_app_instance_bot", Service: "chime", Resource: "app-instance-bot", Template: "arn:${Partition}:chime:${Region}:${AccountId}:app-instance/${AppInstanceId}/bot/${AppInstanceBotId}"},
		{Name: "chime_app_instance_user", Service: "chime", Resource: "app-instance-user", Template: "arn:${Partition}:chime:${Region}:${AccountId}:app-instance/${AppInstanceId}/user/${AppInstanceUserId}"},
		{Name: "chime_channel", Service: "chime", Resource: "channel", Template: "arn:${Partition}:chime:${Region}:${AccountId}:app-instance/${AppInstanceId}/channel/${ChannelId}"},
		{Name: "chime_channel_flow", Service: "chime", Resource: "channel-flow", Template: "arn:${Partition}:chime:${Region}:${AccountId}:app-instance/${AppInstanceId}/channel-flow/${ChannelFlowId}"},
		{Name: "chime_media_insights_pipeline_configuration", Service: "chime", Resource: "media-insights-pipeline-configuration", Template: "arn:${Partition}:chime:${Region}:${AccountId}:media-insights-pipeline-configuration/${ConfigurationName}"},
		{Name: "chime_media_pipeline", Service: "chime", Resource: "media-pipeline", Template: "arn:${Partition}:chime:${Region}:${AccountId}:media-pipeline/${MediaPipelineId}"},
		{Name: "chime_media_pipeline_kinesis_video_stream_pool", Service: "chime", Resource: "media-pipeline-kinesis-video-stream-pool", Template: "arn:${Partition}:chime:${Region}:${AccountId}:media-pipeline-kinesis-video-stream-pool/${PoolName}"},
		{Name: "chime_meeting", Service: "chime", Resource: "meeting", Template: "arn:${Partition}:chime:${Region}:${AccountId}:meeting/${MeetingId}"},
		{Name: "chime_sip_media_application", Service: "chime", Resource: "sip-media-application", Template: "arn:${Partition}:chime:${Region}:${AccountId}:sma/${SipMediaApplicationId}"},
		{Name: "chime_voice_connector", Service: "chime", Resource: "voice-connector", Template: "arn:${Partition}:chime:${Region}:${AccountId}:vc/${VoiceConnectorId}"},
		{Name: "chime_voice_profile", Service: "chime", Resource: "voice-profile", Template: "arn:${Partition}:chime:${Region}:${AccountId}:voice-profile/${VoiceProfileId}"},
		{Name: "chime_voice_profile_domain", Service: "chime", Resource: "voice-profile-domain", Template: "arn:${Partition}:chime:${Region}:${AccountId}:voice-profile-domain/${VoiceProfileDomainId}"},
	})
}
