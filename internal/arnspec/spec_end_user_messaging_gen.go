// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: end-user-messaging
// Source: https://servicereference.us-east-1.amazonaws.com/v1/end-user-messaging/end-user-messaging.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "end_user_messaging_brand_profile", Service: "end-user-messaging", Resource: "brand-profile", Template: "arn:${Partition}:end-user-messaging:${Region}:${Account}:brand-profile/${ResourceId}"},
		{Name: "end_user_messaging_notify_code_configuration", Service: "end-user-messaging", Resource: "notify-code-configuration", Template: "arn:${Partition}:end-user-messaging:${Region}:${Account}:notify-code-configuration/${ResourceId}"},
	})
}
