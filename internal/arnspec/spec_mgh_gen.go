// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: mgh
// Source: https://servicereference.us-east-1.amazonaws.com/v1/mgh/mgh.json
// Functions: 5
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "mgh_automation_run_resource", Service: "mgh", Resource: "AutomationRunResource", Template: "arn:${Partition}:mgh:${Region}:${Account}:automation-run/${RunID}"},
		{Name: "mgh_automation_unit_resource", Service: "mgh", Resource: "AutomationUnitResource", Template: "arn:${Partition}:mgh:${Region}:${Account}:automation-unit/${AutomationUnitId}"},
		{Name: "mgh_connection_resource", Service: "mgh", Resource: "ConnectionResource", Template: "arn:${Partition}:mgh:${Region}:${Account}:${ConnectionArn}"},
		{Name: "mgh_migration_task", Service: "mgh", Resource: "migrationTask", Template: "arn:${Partition}:mgh:${Region}:${Account}:progressUpdateStream/${Stream}/migrationTask/${Task}"},
		{Name: "mgh_progress_update_stream", Service: "mgh", Resource: "progressUpdateStream", Template: "arn:${Partition}:mgh:${Region}:${Account}:progressUpdateStream/${Stream}"},
	})
}
