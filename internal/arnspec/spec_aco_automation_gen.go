// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: aco-automation
// Source: https://servicereference.us-east-1.amazonaws.com/v1/aco-automation/aco-automation.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "aco_automation_automation_rule", Service: "aco-automation", Resource: "AutomationRule", Template: "arn:${Partition}:compute-optimizer::${Account}:automation-rule/${RuleId}"},
	})
}
