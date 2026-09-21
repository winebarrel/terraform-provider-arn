// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: resource-explorer-2
// Source: https://servicereference.us-east-1.amazonaws.com/v1/resource-explorer-2/resource-explorer-2.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "resource_explorer_2_index", Service: "resource-explorer-2", Resource: "index", Template: "arn:${Partition}:resource-explorer-2:${Region}:${Account}:index/${IndexUuid}"},
		{Name: "resource_explorer_2_managed_view", Service: "resource-explorer-2", Resource: "managed-view", Template: "arn:${Partition}:resource-explorer-2:${Region}:${Account}:managed-view/${ManagedViewName}/${ManagedViewUuid}"},
		{Name: "resource_explorer_2_view", Service: "resource-explorer-2", Resource: "view", Template: "arn:${Partition}:resource-explorer-2:${Region}:${Account}:view/${ViewName}/${ViewUuid}"},
	})
}
