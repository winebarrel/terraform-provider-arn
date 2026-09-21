// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: scheduler
// Source: https://servicereference.us-east-1.amazonaws.com/v1/scheduler/scheduler.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "scheduler_schedule", Service: "scheduler", Resource: "schedule", Template: "arn:${Partition}:scheduler:${Region}:${Account}:schedule/${GroupName}/${ScheduleName}"},
		{Name: "scheduler_schedule_group", Service: "scheduler", Resource: "schedule-group", Template: "arn:${Partition}:scheduler:${Region}:${Account}:schedule-group/${GroupName}"},
	})
}
