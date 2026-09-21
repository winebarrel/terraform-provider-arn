// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: qbusiness
// Source: https://servicereference.us-east-1.amazonaws.com/v1/qbusiness/qbusiness.json
// Functions: 10
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "qbusiness_application", Service: "qbusiness", Resource: "application", Template: "arn:${Partition}:qbusiness:${Region}:${Account}:application/${ApplicationId}"},
		{Name: "qbusiness_chat_response_configuration", Service: "qbusiness", Resource: "chat-response-configuration", Template: "arn:${Partition}:qbusiness:${Region}:${Account}:application/${ApplicationId}/chat-response-configuration/${ChatResponseConfigurationId}"},
		{Name: "qbusiness_data_accessor", Service: "qbusiness", Resource: "data-accessor", Template: "arn:${Partition}:qbusiness:${Region}:${Account}:application/${ApplicationId}/data-accessor/${DataAccessorId}"},
		{Name: "qbusiness_data_source", Service: "qbusiness", Resource: "data-source", Template: "arn:${Partition}:qbusiness:${Region}:${Account}:application/${ApplicationId}/index/${IndexId}/data-source/${DataSourceId}"},
		{Name: "qbusiness_index", Service: "qbusiness", Resource: "index", Template: "arn:${Partition}:qbusiness:${Region}:${Account}:application/${ApplicationId}/index/${IndexId}"},
		{Name: "qbusiness_integration", Service: "qbusiness", Resource: "integration", Template: "arn:${Partition}:qbusiness:${Region}:${Account}:application/${ApplicationId}/integration/${IntegrationId}"},
		{Name: "qbusiness_plugin", Service: "qbusiness", Resource: "plugin", Template: "arn:${Partition}:qbusiness:${Region}:${Account}:application/${ApplicationId}/plugin/${PluginId}"},
		{Name: "qbusiness_retriever", Service: "qbusiness", Resource: "retriever", Template: "arn:${Partition}:qbusiness:${Region}:${Account}:application/${ApplicationId}/retriever/${RetrieverId}"},
		{Name: "qbusiness_subscription", Service: "qbusiness", Resource: "subscription", Template: "arn:${Partition}:qbusiness:${Region}:${Account}:application/${ApplicationId}/subscription/${SubscriptionId}"},
		{Name: "qbusiness_web_experience", Service: "qbusiness", Resource: "web-experience", Template: "arn:${Partition}:qbusiness:${Region}:${Account}:application/${ApplicationId}/web-experience/${WebExperienceId}"},
	})
}
