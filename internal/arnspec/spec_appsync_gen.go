// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: appsync
// Source: https://servicereference.us-east-1.amazonaws.com/v1/appsync/appsync.json
// Functions: 10
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "appsync_api", Service: "appsync", Resource: "api", Template: "arn:${Partition}:appsync:${Region}:${Account}:apis/${ApiId}"},
		{Name: "appsync_channel_namespace", Service: "appsync", Resource: "channelNamespace", Template: "arn:${Partition}:appsync:${Region}:${Account}:apis/${ApiId}/channelNamespace/${ChannelNamespaceName}"},
		{Name: "appsync_datasource", Service: "appsync", Resource: "datasource", Template: "arn:${Partition}:appsync:${Region}:${Account}:apis/${GraphQLAPIId}/datasources/${DatasourceName}"},
		{Name: "appsync_domain", Service: "appsync", Resource: "domain", Template: "arn:${Partition}:appsync:${Region}:${Account}:domainnames/${DomainName}"},
		{Name: "appsync_field", Service: "appsync", Resource: "field", Template: "arn:${Partition}:appsync:${Region}:${Account}:apis/${GraphQLAPIId}/types/${TypeName}/fields/${FieldName}"},
		{Name: "appsync_function", Service: "appsync", Resource: "function", Template: "arn:${Partition}:appsync:${Region}:${Account}:apis/${GraphQLAPIId}/functions/${FunctionId}"},
		{Name: "appsync_graphqlapi", Service: "appsync", Resource: "graphqlapi", Template: "arn:${Partition}:appsync:${Region}:${Account}:apis/${GraphQLAPIId}"},
		{Name: "appsync_merged_api_association", Service: "appsync", Resource: "mergedApiAssociation", Template: "arn:${Partition}:appsync:${Region}:${Account}:apis/${SourceGraphQLAPIId}/mergedApiAssociations/${Associationid}"},
		{Name: "appsync_source_api_association", Service: "appsync", Resource: "sourceApiAssociation", Template: "arn:${Partition}:appsync:${Region}:${Account}:apis/${MergedGraphQLAPIId}/sourceApiAssociations/${Associationid}"},
		{Name: "appsync_type", Service: "appsync", Resource: "type", Template: "arn:${Partition}:appsync:${Region}:${Account}:apis/${GraphQLAPIId}/types/${TypeName}"},
	})
}
