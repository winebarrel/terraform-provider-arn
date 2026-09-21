// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: inspector2
// Source: https://servicereference.us-east-1.amazonaws.com/v1/inspector2/inspector2.json
// Functions: 6
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "inspector2_cis_scan_configuration", Service: "inspector2", Resource: "CIS Scan Configuration", Template: "arn:${Partition}:inspector2:${Region}:${Account}:owner/${OwnerId}/cis-configuration/${CISScanConfigurationId}"},
		{Name: "inspector2_code_security_integration", Service: "inspector2", Resource: "Code Security Integration", Template: "arn:${Partition}:inspector2:${Region}:${Account}:codesecurity-integration/${CodeSecurityIntegrationId}"},
		{Name: "inspector2_code_security_scan_configuration", Service: "inspector2", Resource: "Code Security Scan Configuration", Template: "arn:${Partition}:inspector2:${Region}:${Account}:owner/${OwnerId}/codesecurity-configuration/${CodeSecurityScanConfigurationId}"},
		{Name: "inspector2_connector", Service: "inspector2", Resource: "Connector", Template: "arn:${Partition}:inspector2:${Region}:${Account}:connector/${ConnectorId}"},
		{Name: "inspector2_filter", Service: "inspector2", Resource: "Filter", Template: "arn:${Partition}:inspector2:${Region}:${Account}:owner/${OwnerId}/filter/${FilterId}"},
		{Name: "inspector2_finding", Service: "inspector2", Resource: "Finding", Template: "arn:${Partition}:inspector2:${Region}:${Account}:finding/${FindingId}"},
	})
}
