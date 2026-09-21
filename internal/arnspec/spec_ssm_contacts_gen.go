// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: ssm-contacts
// Source: https://servicereference.us-east-1.amazonaws.com/v1/ssm-contacts/ssm-contacts.json
// Functions: 5
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "ssm_contacts_contact", Service: "ssm-contacts", Resource: "contact", Template: "arn:${Partition}:ssm-contacts:${Region}:${Account}:contact/${ContactAlias}"},
		{Name: "ssm_contacts_contactchannel", Service: "ssm-contacts", Resource: "contactchannel", Template: "arn:${Partition}:ssm-contacts:${Region}:${Account}:contactchannel/${ContactAlias}/${ContactChannelId}"},
		{Name: "ssm_contacts_engagement", Service: "ssm-contacts", Resource: "engagement", Template: "arn:${Partition}:ssm-contacts:${Region}:${Account}:engagement/${ContactAlias}/${EngagementId}"},
		{Name: "ssm_contacts_page", Service: "ssm-contacts", Resource: "page", Template: "arn:${Partition}:ssm-contacts:${Region}:${Account}:page/${ContactAlias}/${PageId}"},
		{Name: "ssm_contacts_rotation", Service: "ssm-contacts", Resource: "rotation", Template: "arn:${Partition}:ssm-contacts:${Region}:${Account}:rotation/${RotationId}"},
	})
}
