// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: config
// Source: https://servicereference.us-east-1.amazonaws.com/v1/config/config.json
// Functions: 10
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "config_aggregation_authorization", Service: "config", Resource: "AggregationAuthorization", Template: "arn:${Partition}:config:${Region}:${Account}:aggregation-authorization/${AggregatorAccount}/${AggregatorRegion}"},
		{Name: "config_config_rule", Service: "config", Resource: "ConfigRule", Template: "arn:${Partition}:config:${Region}:${Account}:config-rule/${ConfigRuleId}"},
		{Name: "config_configuration_aggregator", Service: "config", Resource: "ConfigurationAggregator", Template: "arn:${Partition}:config:${Region}:${Account}:config-aggregator/${AggregatorId}"},
		{Name: "config_configuration_recorder", Service: "config", Resource: "ConfigurationRecorder", Template: "arn:${Partition}:config:${Region}:${Account}:configuration-recorder/${RecorderName}/${RecorderId}"},
		{Name: "config_conformance_pack", Service: "config", Resource: "ConformancePack", Template: "arn:${Partition}:config:${Region}:${Account}:conformance-pack/${ConformancePackName}/${ConformancePackId}"},
		{Name: "config_connector", Service: "config", Resource: "Connector", Template: "arn:${Partition}:config:${Region}:${Account}:connector/${Provider}/${ProviderId}/${ConnectorId}"},
		{Name: "config_organization_config_rule", Service: "config", Resource: "OrganizationConfigRule", Template: "arn:${Partition}:config:${Region}:${Account}:organization-config-rule/${OrganizationConfigRuleId}"},
		{Name: "config_organization_conformance_pack", Service: "config", Resource: "OrganizationConformancePack", Template: "arn:${Partition}:config:${Region}:${Account}:organization-conformance-pack/${OrganizationConformancePackId}"},
		{Name: "config_remediation_configuration", Service: "config", Resource: "RemediationConfiguration", Template: "arn:${Partition}:config:${Region}:${Account}:remediation-configuration/${RemediationConfigurationId}"},
		{Name: "config_stored_query", Service: "config", Resource: "StoredQuery", Template: "arn:${Partition}:config:${Region}:${Account}:stored-query/${StoredQueryName}/${StoredQueryId}"},
	})
}
