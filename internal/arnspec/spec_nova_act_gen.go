// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: nova-act
// Source: https://servicereference.us-east-1.amazonaws.com/v1/nova-act/nova-act.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "nova_act_workflow_definition", Service: "nova-act", Resource: "workflow-definition", Template: "arn:${Partition}:nova-act:${Region}:${Account}:workflow-definition/${WorkflowDefinitionName}"},
		{Name: "nova_act_workflow_run", Service: "nova-act", Resource: "workflow-run", Template: "arn:${Partition}:nova-act:${Region}:${Account}:workflow-definition/${WorkflowDefinitionName}/workflow-run/${WorkflowRunId}"},
	})
}
