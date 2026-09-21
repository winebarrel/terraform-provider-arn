// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: entityresolution
// Source: https://servicereference.us-east-1.amazonaws.com/v1/entityresolution/entityresolution.json
// Functions: 5
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "entityresolution_id_mapping_workflow", Service: "entityresolution", Resource: "IdMappingWorkflow", Template: "arn:${Partition}:entityresolution:${Region}:${Account}:idmappingworkflow/${WorkflowName}"},
		{Name: "entityresolution_id_namespace", Service: "entityresolution", Resource: "IdNamespace", Template: "arn:${Partition}:entityresolution:${Region}:${Account}:idnamespace/${IdNamespaceName}"},
		{Name: "entityresolution_matching_workflow", Service: "entityresolution", Resource: "MatchingWorkflow", Template: "arn:${Partition}:entityresolution:${Region}:${Account}:matchingworkflow/${WorkflowName}"},
		{Name: "entityresolution_provider_service", Service: "entityresolution", Resource: "ProviderService", Template: "arn:${Partition}:entityresolution:${Region}:${Account}:providerservice/${ProviderName}/${ProviderServiceName}"},
		{Name: "entityresolution_schema_mapping", Service: "entityresolution", Resource: "SchemaMapping", Template: "arn:${Partition}:entityresolution:${Region}:${Account}:schemamapping/${SchemaName}"},
	})
}
