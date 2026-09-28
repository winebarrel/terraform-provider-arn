// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: network-security-manager
// Source: https://servicereference.us-east-1.amazonaws.com/v1/network-security-manager/network-security-manager.json
// Functions: 10
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "network_security_manager_deployment", Service: "network-security-manager", Resource: "deployment", Template: "arn:${Partition}:network-security-manager:${Region}:${Account}:deployment:${DeploymentId}"},
		{Name: "network_security_manager_deployment_snapshot", Service: "network-security-manager", Resource: "deployment-snapshot", Template: "arn:${Partition}:network-security-manager:${Region}:${Account}:deployment:${DeploymentId}:${VersionNumber}"},
		{Name: "network_security_manager_policy", Service: "network-security-manager", Resource: "policy", Template: "arn:${Partition}:network-security-manager:${Region}:${Account}:policy:${PolicyId}"},
		{Name: "network_security_manager_policy_snapshot", Service: "network-security-manager", Resource: "policy-snapshot", Template: "arn:${Partition}:network-security-manager:${Region}:${Account}:policy:${PolicyId}:${VersionNumber}"},
		{Name: "network_security_manager_rule", Service: "network-security-manager", Resource: "rule", Template: "arn:${Partition}:network-security-manager:${Region}:${Account}:rule:${RuleId}"},
		{Name: "network_security_manager_rule_snapshot", Service: "network-security-manager", Resource: "rule-snapshot", Template: "arn:${Partition}:network-security-manager:${Region}:${Account}:rule:${RuleId}:${VersionNumber}"},
		{Name: "network_security_manager_scope", Service: "network-security-manager", Resource: "scope", Template: "arn:${Partition}:network-security-manager:${Region}:${Account}:scope:${ScopeId}"},
		{Name: "network_security_manager_scope_snapshot", Service: "network-security-manager", Resource: "scope-snapshot", Template: "arn:${Partition}:network-security-manager:${Region}:${Account}:scope:${ScopeId}:${VersionNumber}"},
		{Name: "network_security_manager_template", Service: "network-security-manager", Resource: "template", Template: "arn:${Partition}:network-security-manager:${Region}:${Account}:template:${TemplateId}"},
		{Name: "network_security_manager_template_snapshot", Service: "network-security-manager", Resource: "template-snapshot", Template: "arn:${Partition}:network-security-manager:${Region}:${Account}:template:${TemplateId}:${VersionNumber}"},
	})
}
