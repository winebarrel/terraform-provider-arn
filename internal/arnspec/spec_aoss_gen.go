// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: aoss
// Source: https://servicereference.us-east-1.amazonaws.com/v1/aoss/aoss.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "aoss_collection", Service: "aoss", Resource: "Collection", Template: "arn:${Partition}:aoss:${Region}:${Account}:collection/${CollectionId}"},
		{Name: "aoss_collection_group", Service: "aoss", Resource: "CollectionGroup", Template: "arn:${Partition}:aoss:${Region}:${Account}:collection-group/${CollectionGroupId}"},
		{Name: "aoss_dashboards", Service: "aoss", Resource: "Dashboards", Template: "arn:${Partition}:aoss:${Region}:${Account}:dashboards/default"},
	})
}
