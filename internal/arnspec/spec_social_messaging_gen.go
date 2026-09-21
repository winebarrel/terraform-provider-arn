// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: social-messaging
// Source: https://servicereference.us-east-1.amazonaws.com/v1/social-messaging/social-messaging.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "social_messaging_phone_number_id", Service: "social-messaging", Resource: "phone-number-id", Template: "arn:${Partition}:social-messaging:${Region}:${Account}:phone-number-id/${OriginationPhoneNumberId}"},
		{Name: "social_messaging_waba", Service: "social-messaging", Resource: "waba", Template: "arn:${Partition}:social-messaging:${Region}:${Account}:waba/${WabaId}"},
	})
}
