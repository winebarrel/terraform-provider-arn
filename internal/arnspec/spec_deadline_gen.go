// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: deadline
// Source: https://servicereference.us-east-1.amazonaws.com/v1/deadline/deadline.json
// Functions: 9
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "deadline_budget", Service: "deadline", Resource: "budget", Template: "arn:${Partition}:deadline:${Region}:${Account}:farm/${FarmId}/budget/${BudgetId}"},
		{Name: "deadline_farm", Service: "deadline", Resource: "farm", Template: "arn:${Partition}:deadline:${Region}:${Account}:farm/${FarmId}"},
		{Name: "deadline_fleet", Service: "deadline", Resource: "fleet", Template: "arn:${Partition}:deadline:${Region}:${Account}:farm/${FarmId}/fleet/${FleetId}"},
		{Name: "deadline_job", Service: "deadline", Resource: "job", Template: "arn:${Partition}:deadline:${Region}:${Account}:farm/${FarmId}/queue/${QueueId}/job/${JobId}"},
		{Name: "deadline_license_endpoint", Service: "deadline", Resource: "license-endpoint", Template: "arn:${Partition}:deadline:${Region}:${Account}:license-endpoint/${LicenseEndpointId}"},
		{Name: "deadline_monitor", Service: "deadline", Resource: "monitor", Template: "arn:${Partition}:deadline:${Region}:${Account}:monitor/${MonitorId}"},
		{Name: "deadline_queue", Service: "deadline", Resource: "queue", Template: "arn:${Partition}:deadline:${Region}:${Account}:farm/${FarmId}/queue/${QueueId}"},
		{Name: "deadline_volume", Service: "deadline", Resource: "volume", Template: "arn:${Partition}:deadline:${Region}:${Account}:farm/${FarmId}/fleet/${FleetId}/volume/${VolumeId}"},
		{Name: "deadline_worker", Service: "deadline", Resource: "worker", Template: "arn:${Partition}:deadline:${Region}:${Account}:farm/${FarmId}/fleet/${FleetId}/worker/${WorkerId}"},
	})
}
