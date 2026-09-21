// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: connect-campaigns
// Source: https://servicereference.us-east-1.amazonaws.com/v1/connect-campaigns/connect-campaigns.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "connect_campaigns_campaign", Service: "connect-campaigns", Resource: "campaign", Template: "arn:${Partition}:connect-campaigns:${Region}:${Account}:campaign/${CampaignId}"},
	})
}
