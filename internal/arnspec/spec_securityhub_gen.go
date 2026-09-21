// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: securityhub
// Source: https://servicereference.us-east-1.amazonaws.com/v1/securityhub/securityhub.json
// Functions: 10
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "securityhub_aggregatorv2", Service: "securityhub", Resource: "aggregatorv2", Template: "arn:${Partition}:securityhub:${Region}:${Account}:aggregatorv2/${AggregatorV2Id}"},
		{Name: "securityhub_automation_rule", Service: "securityhub", Resource: "automation-rule", Template: "arn:${Partition}:securityhub:${Region}:${Account}:automation-rule/${AutomationRuleId}"},
		{Name: "securityhub_automation_rulev2", Service: "securityhub", Resource: "automation-rulev2", Template: "arn:${Partition}:securityhub:${Region}:${Account}:automation-rulev2/${AutomationRuleV2Id}"},
		{Name: "securityhub_configuration_policy", Service: "securityhub", Resource: "configuration-policy", Template: "arn:${Partition}:securityhub:${Region}:${Account}:configuration-policy/${ConfigurationPolicyId}"},
		{Name: "securityhub_connector", Service: "securityhub", Resource: "connector", Template: "arn:${Partition}:securityhub:${Region}:${Account}:connector/${ConnectorId}"},
		{Name: "securityhub_connectorv2", Service: "securityhub", Resource: "connectorv2", Template: "arn:${Partition}:securityhub:${Region}:${Account}:connectorv2/${ConnectorV2Id}"},
		{Name: "securityhub_finding_aggregator", Service: "securityhub", Resource: "finding-aggregator", Template: "arn:${Partition}:securityhub:${Region}:${Account}:finding-aggregator/${FindingAggregatorId}"},
		{Name: "securityhub_hub", Service: "securityhub", Resource: "hub", Template: "arn:${Partition}:securityhub:${Region}:${Account}:hub/default"},
		{Name: "securityhub_hubv2", Service: "securityhub", Resource: "hubv2", Template: "arn:${Partition}:securityhub:${Region}:${Account}:hubv2/${HubV2Id}"},
		{Name: "securityhub_product", Service: "securityhub", Resource: "product", Template: "arn:${Partition}:securityhub:${Region}:${Account}:product/${Company}/${ProductId}"},
	})
}
