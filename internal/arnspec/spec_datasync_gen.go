// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: datasync
// Source: https://servicereference.us-east-1.amazonaws.com/v1/datasync/datasync.json
// Functions: 6
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "datasync_agent", Service: "datasync", Resource: "agent", Template: "arn:${Partition}:datasync:${Region}:${AccountId}:agent/${AgentId}"},
		{Name: "datasync_discoveryjob", Service: "datasync", Resource: "discoveryjob", Template: "arn:${Partition}:datasync:${Region}:${AccountId}:system/${StorageSystemId}/job/${DiscoveryJobId}"},
		{Name: "datasync_location", Service: "datasync", Resource: "location", Template: "arn:${Partition}:datasync:${Region}:${AccountId}:location/${LocationId}"},
		{Name: "datasync_storagesystem", Service: "datasync", Resource: "storagesystem", Template: "arn:${Partition}:datasync:${Region}:${AccountId}:system/${StorageSystemId}"},
		{Name: "datasync_task", Service: "datasync", Resource: "task", Template: "arn:${Partition}:datasync:${Region}:${AccountId}:task/${TaskId}"},
		{Name: "datasync_taskexecution", Service: "datasync", Resource: "taskexecution", Template: "arn:${Partition}:datasync:${Region}:${AccountId}:task/${TaskId}/execution/${ExecutionId}"},
	})
}
