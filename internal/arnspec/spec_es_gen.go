// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: es
// Source: https://servicereference.us-east-1.amazonaws.com/v1/es/es.json
// Functions: 5
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "es_application", Service: "es", Resource: "application", Template: "arn:${Partition}:opensearch:${Region}:${Account}:application/${AppId}"},
		{Name: "es_datasource", Service: "es", Resource: "datasource", Template: "arn:${Partition}:opensearch:${Region}:${Account}:datasource/${DataSourceName}"},
		{Name: "es_domain", Service: "es", Resource: "domain", Template: "arn:${Partition}:es:${Region}:${Account}:domain/${DomainName}"},
		{Name: "es_es_role", Service: "es", Resource: "es_role", Template: "arn:${Partition}:iam::${Account}:role/aws-service-role/es.amazonaws.com/AWSServiceRoleForAmazonOpenSearchService"},
		{Name: "es_opensearchservice_role", Service: "es", Resource: "opensearchservice_role", Template: "arn:${Partition}:iam::${Account}:role/aws-service-role/opensearchservice.amazonaws.com/AWSServiceRoleForAmazonOpenSearchService"},
	})
}
