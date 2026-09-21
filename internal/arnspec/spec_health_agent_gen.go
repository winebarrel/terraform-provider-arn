// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: health-agent
// Source: https://servicereference.us-east-1.amazonaws.com/v1/health-agent/health-agent.json
// Functions: 6
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "health_agent_agent", Service: "health-agent", Resource: "Agent", Template: "arn:${Partition}:health-agent:${Region}:${Account}:domain/${DomainId}/agent/${AgentId}"},
		{Name: "health_agent_domain", Service: "health-agent", Resource: "Domain", Template: "arn:${Partition}:health-agent:${Region}:${Account}:domain/${DomainId}"},
		{Name: "health_agent_integration", Service: "health-agent", Resource: "Integration", Template: "arn:${Partition}:health-agent:${Region}:${Account}:domain/${DomainId}/integration/${IntegrationId}"},
		{Name: "health_agent_patient_insights_job", Service: "health-agent", Resource: "PatientInsightsJob", Template: "arn:${Partition}:health-agent:${Region}:${Account}:domain/${DomainId}/patient-insights-job/${JobId}"},
		{Name: "health_agent_session", Service: "health-agent", Resource: "Session", Template: "arn:${Partition}:health-agent:${Region}:${Account}:domain/${DomainId}/session/${SessionId}"},
		{Name: "health_agent_subscription", Service: "health-agent", Resource: "Subscription", Template: "arn:${Partition}:health-agent:${Region}:${Account}:domain/${DomainId}/subscription/${SubscriptionId}"},
	})
}
