// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: notifications
// Source: https://servicereference.us-east-1.amazonaws.com/v1/notifications/notifications.json
// Functions: 6
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "notifications_event_rule", Service: "notifications", Resource: "EventRule", Template: "arn:${Partition}:notifications::${Account}:configuration/${NotificationConfigurationId}/rule/${EventRuleId}"},
		{Name: "notifications_managed_notification_child_event", Service: "notifications", Resource: "ManagedNotificationChildEvent", Template: "arn:${Partition}:notifications::${Account}:managed-notification-configuration/category/${Category}/sub-category/${Subcategory}/event/${NotificationEventId}/child-event/${NotificationChildEventId}"},
		{Name: "notifications_managed_notification_configuration", Service: "notifications", Resource: "ManagedNotificationConfiguration", Template: "arn:${Partition}:notifications::${Account}:managed-notification-configuration/category/${Category}/sub-category/${Subcategory}"},
		{Name: "notifications_managed_notification_event", Service: "notifications", Resource: "ManagedNotificationEvent", Template: "arn:${Partition}:notifications::${Account}:managed-notification-configuration/category/${Category}/sub-category/${Subcategory}/event/${NotificationEventId}"},
		{Name: "notifications_notification_configuration", Service: "notifications", Resource: "NotificationConfiguration", Template: "arn:${Partition}:notifications::${Account}:configuration/${NotificationConfigurationId}"},
		{Name: "notifications_notification_event", Service: "notifications", Resource: "NotificationEvent", Template: "arn:${Partition}:notifications:${Region}:${Account}:configuration/${NotificationConfigurationId}/event/${NotificationEventId}"},
	})
}
