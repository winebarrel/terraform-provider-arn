// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: dataexchange
// Source: https://servicereference.us-east-1.amazonaws.com/v1/dataexchange/dataexchange.json
// Functions: 9
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "dataexchange_assets", Service: "dataexchange", Resource: "assets", Template: "arn:${Partition}:dataexchange:${Region}:${Account}:data-sets/${DataSetId}/revisions/${RevisionId}/assets/${AssetId}"},
		{Name: "dataexchange_data_grants", Service: "dataexchange", Resource: "data-grants", Template: "arn:${Partition}:dataexchange:${Region}:${Account}:data-grants/${DataGrantId}"},
		{Name: "dataexchange_data_sets", Service: "dataexchange", Resource: "data-sets", Template: "arn:${Partition}:dataexchange:${Region}:${Account}:data-sets/${DataSetId}"},
		{Name: "dataexchange_entitled_assets", Service: "dataexchange", Resource: "entitled-assets", Template: "arn:${Partition}:dataexchange:${Region}::data-sets/${DataSetId}/revisions/${RevisionId}/assets/${AssetId}"},
		{Name: "dataexchange_entitled_data_sets", Service: "dataexchange", Resource: "entitled-data-sets", Template: "arn:${Partition}:dataexchange:${Region}::data-sets/${DataSetId}"},
		{Name: "dataexchange_entitled_revisions", Service: "dataexchange", Resource: "entitled-revisions", Template: "arn:${Partition}:dataexchange:${Region}::data-sets/${DataSetId}/revisions/${RevisionId}"},
		{Name: "dataexchange_event_actions", Service: "dataexchange", Resource: "event-actions", Template: "arn:${Partition}:dataexchange:${Region}:${Account}:event-actions/${EventActionId}"},
		{Name: "dataexchange_jobs", Service: "dataexchange", Resource: "jobs", Template: "arn:${Partition}:dataexchange:${Region}:${Account}:jobs/${JobId}"},
		{Name: "dataexchange_revisions", Service: "dataexchange", Resource: "revisions", Template: "arn:${Partition}:dataexchange:${Region}:${Account}:data-sets/${DataSetId}/revisions/${RevisionId}"},
	})
}
