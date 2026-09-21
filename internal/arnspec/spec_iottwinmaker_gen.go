// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: iottwinmaker
// Source: https://servicereference.us-east-1.amazonaws.com/v1/iottwinmaker/iottwinmaker.json
// Functions: 6
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "iottwinmaker_component_type", Service: "iottwinmaker", Resource: "componentType", Template: "arn:${Partition}:iottwinmaker:${Region}:${Account}:workspace/${WorkspaceId}/component-type/${ComponentTypeId}"},
		{Name: "iottwinmaker_entity", Service: "iottwinmaker", Resource: "entity", Template: "arn:${Partition}:iottwinmaker:${Region}:${Account}:workspace/${WorkspaceId}/entity/${EntityId}"},
		{Name: "iottwinmaker_metadata_transfer_job", Service: "iottwinmaker", Resource: "metadataTransferJob", Template: "arn:${Partition}:iottwinmaker:${Region}:${Account}:metadata-transfer-job/${MetadataTransferJobId}"},
		{Name: "iottwinmaker_scene", Service: "iottwinmaker", Resource: "scene", Template: "arn:${Partition}:iottwinmaker:${Region}:${Account}:workspace/${WorkspaceId}/scene/${SceneId}"},
		{Name: "iottwinmaker_sync_job", Service: "iottwinmaker", Resource: "syncJob", Template: "arn:${Partition}:iottwinmaker:${Region}:${Account}:workspace/${WorkspaceId}/sync-job/${SyncJobId}"},
		{Name: "iottwinmaker_workspace", Service: "iottwinmaker", Resource: "workspace", Template: "arn:${Partition}:iottwinmaker:${Region}:${Account}:workspace/${WorkspaceId}"},
	})
}
