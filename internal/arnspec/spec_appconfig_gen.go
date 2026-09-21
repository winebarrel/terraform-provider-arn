// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: appconfig
// Source: https://servicereference.us-east-1.amazonaws.com/v1/appconfig/appconfig.json
// Functions: 11
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "appconfig_application", Service: "appconfig", Resource: "application", Template: "arn:${Partition}:appconfig:${Region}:${Account}:application/${ApplicationId}"},
		{Name: "appconfig_configuration", Service: "appconfig", Resource: "configuration", Template: "arn:${Partition}:appconfig:${Region}:${Account}:application/${ApplicationId}/environment/${EnvironmentId}/configuration/${ConfigurationProfileId}"},
		{Name: "appconfig_configurationprofile", Service: "appconfig", Resource: "configurationprofile", Template: "arn:${Partition}:appconfig:${Region}:${Account}:application/${ApplicationId}/configurationprofile/${ConfigurationProfileId}"},
		{Name: "appconfig_deployment", Service: "appconfig", Resource: "deployment", Template: "arn:${Partition}:appconfig:${Region}:${Account}:application/${ApplicationId}/environment/${EnvironmentId}/deployment/${DeploymentNumber}"},
		{Name: "appconfig_deploymentstrategy", Service: "appconfig", Resource: "deploymentstrategy", Template: "arn:${Partition}:appconfig:${Region}:${Account}:deploymentstrategy/${DeploymentStrategyId}"},
		{Name: "appconfig_environment", Service: "appconfig", Resource: "environment", Template: "arn:${Partition}:appconfig:${Region}:${Account}:application/${ApplicationId}/environment/${EnvironmentId}"},
		{Name: "appconfig_experimentdefinition", Service: "appconfig", Resource: "experimentdefinition", Template: "arn:${Partition}:appconfig:${Region}:${Account}:application/${ApplicationId}/experimentdefinition/${ExperimentDefinitionId}"},
		{Name: "appconfig_experimentrun", Service: "appconfig", Resource: "experimentrun", Template: "arn:${Partition}:appconfig:${Region}:${Account}:application/${ApplicationId}/experimentdefinition/${ExperimentDefinitionId}/experimentrun/${ExperimentRunNumber}"},
		{Name: "appconfig_extension", Service: "appconfig", Resource: "extension", Template: "arn:${Partition}:appconfig:${Region}:${Account}:extension/${ExtensionId}/${ExtensionVersionNumber}"},
		{Name: "appconfig_extensionassociation", Service: "appconfig", Resource: "extensionassociation", Template: "arn:${Partition}:appconfig:${Region}:${Account}:extensionassociation/${ExtensionAssociationId}"},
		{Name: "appconfig_hostedconfigurationversion", Service: "appconfig", Resource: "hostedconfigurationversion", Template: "arn:${Partition}:appconfig:${Region}:${Account}:application/${ApplicationId}/configurationprofile/${ConfigurationProfileId}"},
	})
}
