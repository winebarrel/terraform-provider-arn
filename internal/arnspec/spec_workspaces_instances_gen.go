// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: workspaces-instances
// Source: https://servicereference.us-east-1.amazonaws.com/v1/workspaces-instances/workspaces-instances.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "workspaces_instances_volume_id", Service: "workspaces-instances", Resource: "VolumeId", Template: "arn:${Partition}:ec2:${Region}:${Account}:volume/${VolumeId}"},
		{Name: "workspaces_instances_workspace_instance_id", Service: "workspaces-instances", Resource: "WorkspaceInstanceId", Template: "arn:${Partition}:workspaces-instances:${Region}:${Account}:workspaceinstance/${WorkspaceInstanceId}"},
	})
}
