// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: resource-groups
// Source: https://servicereference.us-east-1.amazonaws.com/v1/resource-groups/resource-groups.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "resource_groups_group", Service: "resource-groups", Resource: "group", Template: "arn:${Partition}:resource-groups:${Region}:${Account}:group/${GroupName}"},
		{Name: "resource_groups_tag_sync_task", Service: "resource-groups", Resource: "tagSyncTask", Template: "arn:${Partition}:resource-groups:${Region}:${Account}:group/${GroupName}/tag-sync-task/${TaskId}"},
	})
}
