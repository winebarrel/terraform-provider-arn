// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: a4b
// Source: https://servicereference.us-east-1.amazonaws.com/v1/a4b/a4b.json
// Functions: 12
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "a4b_addressbook", Service: "a4b", Resource: "addressbook", Template: "arn:${Partition}:a4b:${Region}:${Account}:address-book/${ResourceId}"},
		{Name: "a4b_conferenceprovider", Service: "a4b", Resource: "conferenceprovider", Template: "arn:${Partition}:a4b:${Region}:${Account}:conference-provider/${ResourceId}"},
		{Name: "a4b_contact", Service: "a4b", Resource: "contact", Template: "arn:${Partition}:a4b:${Region}:${Account}:contact/${ResourceId}"},
		{Name: "a4b_device", Service: "a4b", Resource: "device", Template: "arn:${Partition}:a4b:${Region}:${Account}:device/${ResourceId}"},
		{Name: "a4b_gateway", Service: "a4b", Resource: "gateway", Template: "arn:${Partition}:a4b:${Region}:${Account}:gateway/${ResourceId}"},
		{Name: "a4b_gatewaygroup", Service: "a4b", Resource: "gatewaygroup", Template: "arn:${Partition}:a4b:${Region}:${Account}:gateway-group/${ResourceId}"},
		{Name: "a4b_networkprofile", Service: "a4b", Resource: "networkprofile", Template: "arn:${Partition}:a4b:${Region}:${Account}:network-profile/${ResourceId}"},
		{Name: "a4b_profile", Service: "a4b", Resource: "profile", Template: "arn:${Partition}:a4b:${Region}:${Account}:profile/${ResourceId}"},
		{Name: "a4b_room", Service: "a4b", Resource: "room", Template: "arn:${Partition}:a4b:${Region}:${Account}:room/${ResourceId}"},
		{Name: "a4b_schedule", Service: "a4b", Resource: "schedule", Template: "arn:${Partition}:a4b:${Region}:${Account}:schedule/${ResourceId}"},
		{Name: "a4b_skillgroup", Service: "a4b", Resource: "skillgroup", Template: "arn:${Partition}:a4b:${Region}:${Account}:skill-group/${ResourceId}"},
		{Name: "a4b_user", Service: "a4b", Resource: "user", Template: "arn:${Partition}:a4b:${Region}:${Account}:user/${ResourceId}"},
	})
}
