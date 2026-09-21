// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: securityagent
// Source: https://servicereference.us-east-1.amazonaws.com/v1/securityagent/securityagent.json
// Functions: 6
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "securityagent_agent_space", Service: "securityagent", Resource: "AgentSpace", Template: "arn:${Partition}:securityagent:${Region}:${Account}:agent-space/${AgentId}"},
		{Name: "securityagent_application", Service: "securityagent", Resource: "Application", Template: "arn:${Partition}:securityagent:${Region}:${Account}:application/${ApplicationId}"},
		{Name: "securityagent_integration", Service: "securityagent", Resource: "Integration", Template: "arn:${Partition}:securityagent:${Region}:${Account}:integration/${IntegrationId}"},
		{Name: "securityagent_private_connection", Service: "securityagent", Resource: "PrivateConnection", Template: "arn:${Partition}:securityagent:${Region}:${Account}:private-connection/${PrivateConnectionName}"},
		{Name: "securityagent_security_requirement_pack", Service: "securityagent", Resource: "SecurityRequirementPack", Template: "arn:${Partition}:securityagent:${Region}:${Account}:security-requirement-pack/${SecurityRequirementPackId}"},
		{Name: "securityagent_target_domain", Service: "securityagent", Resource: "TargetDomain", Template: "arn:${Partition}:securityagent:${Region}:${Account}:target-domain/${TargetDomainId}"},
	})
}
