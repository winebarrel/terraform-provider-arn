// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: cloudtrail
// Source: https://servicereference.us-east-1.amazonaws.com/v1/cloudtrail/cloudtrail.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "cloudtrail_channel", Service: "cloudtrail", Resource: "channel", Template: "arn:${Partition}:cloudtrail:${Region}:${Account}:channel/${ChannelId}"},
		{Name: "cloudtrail_dashboard", Service: "cloudtrail", Resource: "dashboard", Template: "arn:${Partition}:cloudtrail:${Region}:${Account}:dashboard/${DashboardName}"},
		{Name: "cloudtrail_eventdatastore", Service: "cloudtrail", Resource: "eventdatastore", Template: "arn:${Partition}:cloudtrail:${Region}:${Account}:eventdatastore/${EventDataStoreId}"},
		{Name: "cloudtrail_trail", Service: "cloudtrail", Resource: "trail", Template: "arn:${Partition}:cloudtrail:${Region}:${Account}:trail/${TrailName}"},
	})
}
