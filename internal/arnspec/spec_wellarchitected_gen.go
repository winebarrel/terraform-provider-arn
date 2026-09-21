// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: wellarchitected
// Source: https://servicereference.us-east-1.amazonaws.com/v1/wellarchitected/wellarchitected.json
// Functions: 6
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "wellarchitected_agent_profile", Service: "wellarchitected", Resource: "agent-profile", Template: "arn:${Partition}:wellarchitected:${Region}:${Account}:agent-profile/${ProfileName}"},
		{Name: "wellarchitected_agent_recommendation", Service: "wellarchitected", Resource: "agent-recommendation", Template: "arn:${Partition}:wellarchitected:${Region}:${Account}:agent-recommendation/${ResourceId}"},
		{Name: "wellarchitected_lens", Service: "wellarchitected", Resource: "lens", Template: "arn:${Partition}:wellarchitected:${Region}:${Account}:lens/${ResourceId}"},
		{Name: "wellarchitected_profile", Service: "wellarchitected", Resource: "profile", Template: "arn:${Partition}:wellarchitected:${Region}:${Account}:profile/${ResourceId}"},
		{Name: "wellarchitected_review_template", Service: "wellarchitected", Resource: "review-template", Template: "arn:${Partition}:wellarchitected:${Region}:${Account}:review-template/${ResourceId}"},
		{Name: "wellarchitected_workload", Service: "wellarchitected", Resource: "workload", Template: "arn:${Partition}:wellarchitected:${Region}:${Account}:workload/${ResourceId}"},
	})
}
