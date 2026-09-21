// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: notifications-contacts
// Source: https://servicereference.us-east-1.amazonaws.com/v1/notifications-contacts/notifications-contacts.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "notifications_contacts_email_contact_resource", Service: "notifications-contacts", Resource: "EmailContactResource", Template: "arn:${Partition}:notifications-contacts::${Account}:emailcontact/${EmailContactId}"},
	})
}
