// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: aidevops
// Source: https://servicereference.us-east-1.amazonaws.com/v1/aidevops/aidevops.json
// Functions: 6
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "aidevops_agentspace", Service: "aidevops", Resource: "agentspace", Template: "arn:${Partition}:aidevops:${Region}:${Account}:agentspace/${AgentSpaceId}"},
		{Name: "aidevops_asset", Service: "aidevops", Resource: "asset", Template: "arn:${Partition}:aidevops:${Region}:${Account}:agentspace/${AgentSpaceId}/asset/${AssetId}"},
		{Name: "aidevops_associations", Service: "aidevops", Resource: "associations", Template: "arn:${Partition}:aidevops:${Region}:${Account}:agentspace/${AgentSpaceId}/association/${AssociationId}"},
		{Name: "aidevops_private_connection", Service: "aidevops", Resource: "private-connection", Template: "arn:${Partition}:aidevops:${Region}:${Account}:private-connection/${Name}"},
		{Name: "aidevops_service", Service: "aidevops", Resource: "service", Template: "arn:${Partition}:aidevops:${Region}:${Account}:service/${ServiceId}"},
		{Name: "aidevops_trigger", Service: "aidevops", Resource: "trigger", Template: "arn:${Partition}:aidevops:${Region}:${Account}:agentspace/${AgentSpaceId}/trigger/${TriggerId}"},
	})
}
