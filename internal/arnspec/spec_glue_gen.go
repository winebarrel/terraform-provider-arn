// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: glue
// Source: https://servicereference.us-east-1.amazonaws.com/v1/glue/glue.json
// Functions: 24
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "glue_blueprint", Service: "glue", Resource: "blueprint", Template: "arn:${Partition}:glue:${Region}:${Account}:blueprint/${BlueprintName}"},
		{Name: "glue_catalog", Service: "glue", Resource: "catalog", Template: "arn:${Partition}:glue:${Region}:${Account}:catalog/${CatalogName}"},
		{Name: "glue_completion", Service: "glue", Resource: "completion", Template: "arn:${Partition}:glue:${Region}:${Account}:completion/${CompletionId}"},
		{Name: "glue_connection", Service: "glue", Resource: "connection", Template: "arn:${Partition}:glue:${Region}:${Account}:connection/${ConnectionName}"},
		{Name: "glue_connection_type", Service: "glue", Resource: "connectionType", Template: "arn:${Partition}:glue:${Region}:${Account}:connectionType:${ConnectionTypeName}"},
		{Name: "glue_crawler", Service: "glue", Resource: "crawler", Template: "arn:${Partition}:glue:${Region}:${Account}:crawler/${CrawlerName}"},
		{Name: "glue_custom_entity_type", Service: "glue", Resource: "customEntityType", Template: "arn:${Partition}:glue:${Region}:${Account}:customEntityType/${CustomEntityTypeId}"},
		{Name: "glue_data_quality_ruleset", Service: "glue", Resource: "dataQualityRuleset", Template: "arn:${Partition}:glue:${Region}:${Account}:dataQualityRuleset/${RulesetName}"},
		{Name: "glue_database", Service: "glue", Resource: "database", Template: "arn:${Partition}:glue:${Region}:${Account}:database/${DatabaseName}"},
		{Name: "glue_devendpoint", Service: "glue", Resource: "devendpoint", Template: "arn:${Partition}:glue:${Region}:${Account}:devEndpoint/${DevEndpointName}"},
		{Name: "glue_integration", Service: "glue", Resource: "integration", Template: "arn:${Partition}:glue:${Region}:${Account}:integration:${IntegrationId}"},
		{Name: "glue_integration_resource_property", Service: "glue", Resource: "integrationResourceProperty", Template: "arn:${Partition}:glue:${Region}:${Account}:integrationresourceproperty/${ResourceType}/${ResourceName}"},
		{Name: "glue_job", Service: "glue", Resource: "job", Template: "arn:${Partition}:glue:${Region}:${Account}:job/${JobName}"},
		{Name: "glue_ml_transform", Service: "glue", Resource: "mlTransform", Template: "arn:${Partition}:glue:${Region}:${Account}:mlTransform/${TransformId}"},
		{Name: "glue_registry", Service: "glue", Resource: "registry", Template: "arn:${Partition}:glue:${Region}:${Account}:registry/${RegistryName}"},
		{Name: "glue_rootcatalog", Service: "glue", Resource: "rootcatalog", Template: "arn:${Partition}:glue:${Region}:${Account}:catalog"},
		{Name: "glue_schema", Service: "glue", Resource: "schema", Template: "arn:${Partition}:glue:${Region}:${Account}:schema/${SchemaName}"},
		{Name: "glue_session", Service: "glue", Resource: "session", Template: "arn:${Partition}:glue:${Region}:${Account}:session/${SessionId}"},
		{Name: "glue_table", Service: "glue", Resource: "table", Template: "arn:${Partition}:glue:${Region}:${Account}:table/${DatabaseName}/${TableName}"},
		{Name: "glue_tableversion", Service: "glue", Resource: "tableversion", Template: "arn:${Partition}:glue:${Region}:${Account}:tableVersion/${DatabaseName}/${TableName}/${TableVersionName}"},
		{Name: "glue_trigger", Service: "glue", Resource: "trigger", Template: "arn:${Partition}:glue:${Region}:${Account}:trigger/${TriggerName}"},
		{Name: "glue_usage_profile", Service: "glue", Resource: "usageProfile", Template: "arn:${Partition}:glue:${Region}:${Account}:usageProfile/${UsageProfileId}"},
		{Name: "glue_userdefinedfunction", Service: "glue", Resource: "userdefinedfunction", Template: "arn:${Partition}:glue:${Region}:${Account}:userDefinedFunction/${DatabaseName}/${UserDefinedFunctionName}"},
		{Name: "glue_workflow", Service: "glue", Resource: "workflow", Template: "arn:${Partition}:glue:${Region}:${Account}:workflow/${WorkflowName}"},
	})
}
