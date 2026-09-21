// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: sms-voice
// Source: https://servicereference.us-east-1.amazonaws.com/v1/sms-voice/sms-voice.json
// Functions: 12
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "sms_voice_configuration_set", Service: "sms-voice", Resource: "ConfigurationSet", Template: "arn:${Partition}:sms-voice:${Region}:${Account}:configuration-set/${ConfigurationSetName}"},
		{Name: "sms_voice_message", Service: "sms-voice", Resource: "Message", Template: "arn:${Partition}:sms-voice:${Region}:${Account}:message/${MessageId}"},
		{Name: "sms_voice_notify_configuration", Service: "sms-voice", Resource: "NotifyConfiguration", Template: "arn:${Partition}:sms-voice:${Region}:${Account}:notify-configuration/${NotifyConfigurationId}"},
		{Name: "sms_voice_opt_out_list", Service: "sms-voice", Resource: "OptOutList", Template: "arn:${Partition}:sms-voice:${Region}:${Account}:opt-out-list/${OptOutListName}"},
		{Name: "sms_voice_phone_number", Service: "sms-voice", Resource: "PhoneNumber", Template: "arn:${Partition}:sms-voice:${Region}:${Account}:phone-number/${PhoneNumberId}"},
		{Name: "sms_voice_pool", Service: "sms-voice", Resource: "Pool", Template: "arn:${Partition}:sms-voice:${Region}:${Account}:pool/${PoolId}"},
		{Name: "sms_voice_protect_configuration", Service: "sms-voice", Resource: "ProtectConfiguration", Template: "arn:${Partition}:sms-voice:${Region}:${Account}:protect-configuration/${ProtectConfigurationId}"},
		{Name: "sms_voice_rcs_agent", Service: "sms-voice", Resource: "RcsAgent", Template: "arn:${Partition}:sms-voice:${Region}:${Account}:rcs-agent/${RcsAgentId}"},
		{Name: "sms_voice_registration", Service: "sms-voice", Resource: "Registration", Template: "arn:${Partition}:sms-voice:${Region}:${Account}:registration/${RegistrationId}"},
		{Name: "sms_voice_registration_attachment", Service: "sms-voice", Resource: "RegistrationAttachment", Template: "arn:${Partition}:sms-voice:${Region}:${Account}:registration-attachment/${RegistrationAttachmentId}"},
		{Name: "sms_voice_sender_id", Service: "sms-voice", Resource: "SenderId", Template: "arn:${Partition}:sms-voice:${Region}:${Account}:sender-id/${SenderId}/${IsoCountryCode}"},
		{Name: "sms_voice_verified_destination_number", Service: "sms-voice", Resource: "VerifiedDestinationNumber", Template: "arn:${Partition}:sms-voice:${Region}:${Account}:verified-destination-number/${VerifiedDestinationNumberId}"},
	})
}
