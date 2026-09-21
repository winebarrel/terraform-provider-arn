// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: agent-registry
// Source: https://servicereference.us-east-1.amazonaws.com/v1/agent-registry/agent-registry.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "agent_registry_registry", Service: "agent-registry", Resource: "registry", Template: "arn:${Partition}:agent-registry:${Region}:${Account}:registry/${RegistryId}"},
		{Name: "agent_registry_registry_record", Service: "agent-registry", Resource: "registry-record", Template: "arn:${Partition}:agent-registry:${Region}:${Account}:registry/${RegistryId}/record/${RecordId}"},
	})
}
