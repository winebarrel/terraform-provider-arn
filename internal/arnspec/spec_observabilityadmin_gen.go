// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: observabilityadmin
// Source: https://servicereference.us-east-1.amazonaws.com/v1/observabilityadmin/observabilityadmin.json
// Functions: 5
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "observabilityadmin_organization_centralization_rule", Service: "observabilityadmin", Resource: "organization-centralization-rule", Template: "arn:${Partition}:observabilityadmin:${Region}:${Account}:organization-centralization-rule/${CentralizationRuleName}"},
		{Name: "observabilityadmin_organization_telemetry_rule", Service: "observabilityadmin", Resource: "organization-telemetry-rule", Template: "arn:${Partition}:observabilityadmin:${Region}:${Account}:organization-telemetry-rule/${TelemetryRuleName}"},
		{Name: "observabilityadmin_s3tableintegration", Service: "observabilityadmin", Resource: "s3tableintegration", Template: "arn:${Partition}:observabilityadmin:${Region}:${Account}:s3tableintegration/${S3TableIntegrationIdentifier}"},
		{Name: "observabilityadmin_telemetry_pipeline", Service: "observabilityadmin", Resource: "telemetry-pipeline", Template: "arn:${Partition}:observabilityadmin:${Region}:${Account}:telemetry-pipeline/${TelemetryPipelineIdentifier}"},
		{Name: "observabilityadmin_telemetry_rule", Service: "observabilityadmin", Resource: "telemetry-rule", Template: "arn:${Partition}:observabilityadmin:${Region}:${Account}:telemetry-rule/${TelemetryRuleName}"},
	})
}
