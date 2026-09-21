// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: greengrass
// Source: https://servicereference.us-east-1.amazonaws.com/v1/greengrass/greengrass.json
// Functions: 26
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "greengrass_bulk_deployment", Service: "greengrass", Resource: "bulkDeployment", Template: "arn:${Partition}:greengrass:${Region}:${Account}:/greengrass/bulk/deployments/${BulkDeploymentId}"},
		{Name: "greengrass_certificate_authority", Service: "greengrass", Resource: "certificateAuthority", Template: "arn:${Partition}:greengrass:${Region}:${Account}:/greengrass/groups/${GroupId}/certificateauthorities/${CertificateAuthorityId}"},
		{Name: "greengrass_component", Service: "greengrass", Resource: "component", Template: "arn:${Partition}:greengrass:${Region}:${Account}:components:${ComponentName}"},
		{Name: "greengrass_component_version", Service: "greengrass", Resource: "componentVersion", Template: "arn:${Partition}:greengrass:${Region}:${Account}:components:${ComponentName}:versions:${ComponentVersion}"},
		{Name: "greengrass_connectivity_info", Service: "greengrass", Resource: "connectivityInfo", Template: "arn:${Partition}:greengrass:${Region}:${Account}:/greengrass/things/${ThingName}/connectivityInfo"},
		{Name: "greengrass_connector_definition", Service: "greengrass", Resource: "connectorDefinition", Template: "arn:${Partition}:greengrass:${Region}:${Account}:/greengrass/definition/connectors/${ConnectorDefinitionId}"},
		{Name: "greengrass_connector_definition_version", Service: "greengrass", Resource: "connectorDefinitionVersion", Template: "arn:${Partition}:greengrass:${Region}:${Account}:/greengrass/definition/connectors/${ConnectorDefinitionId}/versions/${VersionId}"},
		{Name: "greengrass_core_definition", Service: "greengrass", Resource: "coreDefinition", Template: "arn:${Partition}:greengrass:${Region}:${Account}:/greengrass/definition/cores/${CoreDefinitionId}"},
		{Name: "greengrass_core_definition_version", Service: "greengrass", Resource: "coreDefinitionVersion", Template: "arn:${Partition}:greengrass:${Region}:${Account}:/greengrass/definition/cores/${CoreDefinitionId}/versions/${VersionId}"},
		{Name: "greengrass_core_device", Service: "greengrass", Resource: "coreDevice", Template: "arn:${Partition}:greengrass:${Region}:${Account}:coreDevices:${CoreDeviceThingName}"},
		{Name: "greengrass_deployment", Service: "greengrass", Resource: "deployment", Template: "arn:${Partition}:greengrass:${Region}:${Account}:/greengrass/groups/${GroupId}/deployments/${DeploymentId}"},
		{Name: "greengrass_deployment_2", Service: "greengrass", Resource: "deployment", Template: "arn:${Partition}:greengrass:${Region}:${Account}:deployments:${DeploymentId}"},
		{Name: "greengrass_device_definition", Service: "greengrass", Resource: "deviceDefinition", Template: "arn:${Partition}:greengrass:${Region}:${Account}:/greengrass/definition/devices/${DeviceDefinitionId}"},
		{Name: "greengrass_device_definition_version", Service: "greengrass", Resource: "deviceDefinitionVersion", Template: "arn:${Partition}:greengrass:${Region}:${Account}:/greengrass/definition/devices/${DeviceDefinitionId}/versions/${VersionId}"},
		{Name: "greengrass_function_definition", Service: "greengrass", Resource: "functionDefinition", Template: "arn:${Partition}:greengrass:${Region}:${Account}:/greengrass/definition/functions/${FunctionDefinitionId}"},
		{Name: "greengrass_function_definition_version", Service: "greengrass", Resource: "functionDefinitionVersion", Template: "arn:${Partition}:greengrass:${Region}:${Account}:/greengrass/definition/functions/${FunctionDefinitionId}/versions/${VersionId}"},
		{Name: "greengrass_group", Service: "greengrass", Resource: "group", Template: "arn:${Partition}:greengrass:${Region}:${Account}:/greengrass/groups/${GroupId}"},
		{Name: "greengrass_group_version", Service: "greengrass", Resource: "groupVersion", Template: "arn:${Partition}:greengrass:${Region}:${Account}:/greengrass/groups/${GroupId}/versions/${VersionId}"},
		{Name: "greengrass_logger_definition", Service: "greengrass", Resource: "loggerDefinition", Template: "arn:${Partition}:greengrass:${Region}:${Account}:/greengrass/definition/loggers/${LoggerDefinitionId}"},
		{Name: "greengrass_logger_definition_version", Service: "greengrass", Resource: "loggerDefinitionVersion", Template: "arn:${Partition}:greengrass:${Region}:${Account}:/greengrass/definition/loggers/${LoggerDefinitionId}/versions/${VersionId}"},
		{Name: "greengrass_resource_definition", Service: "greengrass", Resource: "resourceDefinition", Template: "arn:${Partition}:greengrass:${Region}:${Account}:/greengrass/definition/resources/${ResourceDefinitionId}"},
		{Name: "greengrass_resource_definition_version", Service: "greengrass", Resource: "resourceDefinitionVersion", Template: "arn:${Partition}:greengrass:${Region}:${Account}:/greengrass/definition/resources/${ResourceDefinitionId}/versions/${VersionId}"},
		{Name: "greengrass_subscription_definition", Service: "greengrass", Resource: "subscriptionDefinition", Template: "arn:${Partition}:greengrass:${Region}:${Account}:/greengrass/definition/subscriptions/${SubscriptionDefinitionId}"},
		{Name: "greengrass_subscription_definition_version", Service: "greengrass", Resource: "subscriptionDefinitionVersion", Template: "arn:${Partition}:greengrass:${Region}:${Account}:/greengrass/definition/subscriptions/${SubscriptionDefinitionId}/versions/${VersionId}"},
		{Name: "greengrass_thing", Service: "greengrass", Resource: "thing", Template: "arn:${Partition}:iot:${Region}:${Account}:thing/${ThingName}"},
		{Name: "greengrass_thing_runtime_config", Service: "greengrass", Resource: "thingRuntimeConfig", Template: "arn:${Partition}:greengrass:${Region}:${Account}:/greengrass/things/${ThingName}/runtimeconfig"},
	})
}
