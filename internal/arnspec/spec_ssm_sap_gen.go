// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: ssm-sap
// Source: https://servicereference.us-east-1.amazonaws.com/v1/ssm-sap/ssm-sap.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "ssm_sap_application", Service: "ssm-sap", Resource: "application", Template: "arn:${Partition}:ssm-sap:${Region}:${Account}:${ApplicationType}/${ApplicationId}"},
		{Name: "ssm_sap_component", Service: "ssm-sap", Resource: "component", Template: "arn:${Partition}:ssm-sap:${Region}:${Account}:${ApplicationType}/${ApplicationId}/COMPONENT/${ComponentId}"},
		{Name: "ssm_sap_database", Service: "ssm-sap", Resource: "database", Template: "arn:${Partition}:ssm-sap:${Region}:${Account}:${ApplicationType}/${ApplicationId}/DB/${DatabaseId}"},
	})
}
