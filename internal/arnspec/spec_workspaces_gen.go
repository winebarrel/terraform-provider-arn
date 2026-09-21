// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: workspaces
// Source: https://servicereference.us-east-1.amazonaws.com/v1/workspaces/workspaces.json
// Functions: 9
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "workspaces_certificateid", Service: "workspaces", Resource: "certificateid", Template: "arn:${Partition}:workspaces:${Region}:${Account}:workspacecertificate/${CertificateId}"},
		{Name: "workspaces_connectionalias", Service: "workspaces", Resource: "connectionalias", Template: "arn:${Partition}:workspaces:${Region}:${Account}:connectionalias/${ConnectionAliasId}"},
		{Name: "workspaces_directoryid", Service: "workspaces", Resource: "directoryid", Template: "arn:${Partition}:workspaces:${Region}:${Account}:directory/${DirectoryId}"},
		{Name: "workspaces_workspaceapplication", Service: "workspaces", Resource: "workspaceapplication", Template: "arn:${Partition}:workspaces:${Region}:${Account}:workspaceapplication/${WorkSpaceApplicationId}"},
		{Name: "workspaces_workspacebundle", Service: "workspaces", Resource: "workspacebundle", Template: "arn:${Partition}:workspaces:${Region}:${Account}:workspacebundle/${BundleId}"},
		{Name: "workspaces_workspaceid", Service: "workspaces", Resource: "workspaceid", Template: "arn:${Partition}:workspaces:${Region}:${Account}:workspace/${WorkspaceId}"},
		{Name: "workspaces_workspaceimage", Service: "workspaces", Resource: "workspaceimage", Template: "arn:${Partition}:workspaces:${Region}:${Account}:workspaceimage/${ImageId}"},
		{Name: "workspaces_workspaceipgroup", Service: "workspaces", Resource: "workspaceipgroup", Template: "arn:${Partition}:workspaces:${Region}:${Account}:workspaceipgroup/${GroupId}"},
		{Name: "workspaces_workspacespool", Service: "workspaces", Resource: "workspacespool", Template: "arn:${Partition}:workspaces:${Region}:${Account}:workspacespool/${PoolId}"},
	})
}
