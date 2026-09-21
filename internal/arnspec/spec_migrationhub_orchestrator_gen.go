// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: migrationhub-orchestrator
// Source: https://servicereference.us-east-1.amazonaws.com/v1/migrationhub-orchestrator/migrationhub-orchestrator.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "migrationhub_orchestrator_template", Service: "migrationhub-orchestrator", Resource: "template", Template: "arn:${Partition}:migrationhub-orchestrator:${Region}:${Account}:template/${ResourceId}"},
		{Name: "migrationhub_orchestrator_workflow", Service: "migrationhub-orchestrator", Resource: "workflow", Template: "arn:${Partition}:migrationhub-orchestrator:${Region}:${Account}:workflow/${ResourceId}"},
	})
}
