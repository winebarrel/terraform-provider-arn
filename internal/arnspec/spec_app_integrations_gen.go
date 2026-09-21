// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: app-integrations
// Source: https://servicereference.us-east-1.amazonaws.com/v1/app-integrations/app-integrations.json
// Functions: 6
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "app_integrations_application", Service: "app-integrations", Resource: "application", Template: "arn:${Partition}:app-integrations:${Region}:${Account}:application/${ApplicationId}"},
		{Name: "app_integrations_application_association", Service: "app-integrations", Resource: "application-association", Template: "arn:${Partition}:app-integrations:${Region}:${Account}:application-association/${ApplicationId}/${ApplicationAssociationId}"},
		{Name: "app_integrations_data_integration", Service: "app-integrations", Resource: "data-integration", Template: "arn:${Partition}:app-integrations:${Region}:${Account}:data-integration/${DataIntegrationId}"},
		{Name: "app_integrations_data_integration_association", Service: "app-integrations", Resource: "data-integration-association", Template: "arn:${Partition}:app-integrations:${Region}:${Account}:data-integration-association/${DataIntegrationId}/${ResourceId}"},
		{Name: "app_integrations_event_integration", Service: "app-integrations", Resource: "event-integration", Template: "arn:${Partition}:app-integrations:${Region}:${Account}:event-integration/${EventIntegrationName}"},
		{Name: "app_integrations_event_integration_association", Service: "app-integrations", Resource: "event-integration-association", Template: "arn:${Partition}:app-integrations:${Region}:${Account}:event-integration-association/${EventIntegrationName}/${ResourceId}"},
	})
}
