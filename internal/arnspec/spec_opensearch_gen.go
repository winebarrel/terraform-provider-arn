// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: opensearch
// Source: https://servicereference.us-east-1.amazonaws.com/v1/opensearch/opensearch.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "opensearch_application", Service: "opensearch", Resource: "application", Template: "arn:${Partition}:opensearch:${Region}:${Account}:application/${AppId}"},
		{Name: "opensearch_datasource", Service: "opensearch", Resource: "datasource", Template: "arn:${Partition}:opensearch:${Region}:${Account}:datasource/${DataSourceName}"},
	})
}
