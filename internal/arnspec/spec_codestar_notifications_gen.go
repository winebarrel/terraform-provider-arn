// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: codestar-notifications
// Source: https://servicereference.us-east-1.amazonaws.com/v1/codestar-notifications/codestar-notifications.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "codestar_notifications_notificationrule", Service: "codestar-notifications", Resource: "notificationrule", Template: "arn:${Partition}:codestar-notifications:${Region}:${Account}:notificationrule/${NotificationRuleId}"},
	})
}
