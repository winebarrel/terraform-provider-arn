// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: controltower
// Source: https://servicereference.us-east-1.amazonaws.com/v1/controltower/controltower.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "controltower_baseline", Service: "controltower", Resource: "Baseline", Template: "arn:${Partition}:controltower:${Region}::baseline/${BaselineId}"},
		{Name: "controltower_enabled_baseline", Service: "controltower", Resource: "EnabledBaseline", Template: "arn:${Partition}:controltower:${Region}:${Account}:enabledbaseline/${EnabledBaselineId}"},
		{Name: "controltower_enabled_control", Service: "controltower", Resource: "EnabledControl", Template: "arn:${Partition}:controltower:${Region}:${Account}:enabledcontrol/${EnabledControlId}"},
		{Name: "controltower_landing_zone", Service: "controltower", Resource: "LandingZone", Template: "arn:${Partition}:controltower:${Region}:${Account}:landingzone/${LandingZoneId}"},
	})
}
