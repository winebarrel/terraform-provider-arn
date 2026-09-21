// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: guardduty
// Source: https://servicereference.us-east-1.amazonaws.com/v1/guardduty/guardduty.json
// Functions: 10
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "guardduty_customdetectionrule", Service: "guardduty", Resource: "customdetectionrule", Template: "arn:${Partition}:guardduty::aws:detection-rule/custom/${RuleId}"},
		{Name: "guardduty_customdetectionruleassociation", Service: "guardduty", Resource: "customdetectionruleassociation", Template: "arn:${Partition}:guardduty:${Region}:${Account}:detection-rule/custom/${RuleId}/association/${AssociationId}"},
		{Name: "guardduty_detector", Service: "guardduty", Resource: "detector", Template: "arn:${Partition}:guardduty:${Region}:${Account}:detector/${DetectorId}"},
		{Name: "guardduty_filter", Service: "guardduty", Resource: "filter", Template: "arn:${Partition}:guardduty:${Region}:${Account}:detector/${DetectorId}/filter/${FilterName}"},
		{Name: "guardduty_ipset", Service: "guardduty", Resource: "ipset", Template: "arn:${Partition}:guardduty:${Region}:${Account}:detector/${DetectorId}/ipset/${IPSetId}"},
		{Name: "guardduty_malwareprotectionplan", Service: "guardduty", Resource: "malwareprotectionplan", Template: "arn:${Partition}:guardduty:${Region}:${Account}:malware-protection-plan/${MalwareProtectionPlanId}"},
		{Name: "guardduty_publishing_destination", Service: "guardduty", Resource: "publishingDestination", Template: "arn:${Partition}:guardduty:${Region}:${Account}:detector/${DetectorId}/publishingdestination/${PublishingDestinationId}"},
		{Name: "guardduty_threatentityset", Service: "guardduty", Resource: "threatentityset", Template: "arn:${Partition}:guardduty:${Region}:${Account}:detector/${DetectorId}/threatentityset/${ThreatEntitySetId}"},
		{Name: "guardduty_threatintelset", Service: "guardduty", Resource: "threatintelset", Template: "arn:${Partition}:guardduty:${Region}:${Account}:detector/${DetectorId}/threatintelset/${ThreatIntelSetId}"},
		{Name: "guardduty_trustedentityset", Service: "guardduty", Resource: "trustedentityset", Template: "arn:${Partition}:guardduty:${Region}:${Account}:detector/${DetectorId}/trustedentityset/${TrustedEntitySetId}"},
	})
}
