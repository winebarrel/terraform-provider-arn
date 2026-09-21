// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: kafkaconnect
// Source: https://servicereference.us-east-1.amazonaws.com/v1/kafkaconnect/kafkaconnect.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "kafkaconnect_connector", Service: "kafkaconnect", Resource: "connector", Template: "arn:${Partition}:kafkaconnect:${Region}:${Account}:connector/${ConnectorName}/${UUID}"},
		{Name: "kafkaconnect_connector_operation", Service: "kafkaconnect", Resource: "connector operation", Template: "arn:${Partition}:kafkaconnect:${Region}:${Account}:connector-operation/${ConnectorName}/${ConnectorUUID}/${UUID}"},
		{Name: "kafkaconnect_custom_plugin", Service: "kafkaconnect", Resource: "custom plugin", Template: "arn:${Partition}:kafkaconnect:${Region}:${Account}:custom-plugin/${CustomPluginName}/${UUID}"},
		{Name: "kafkaconnect_worker_configuration", Service: "kafkaconnect", Resource: "worker configuration", Template: "arn:${Partition}:kafkaconnect:${Region}:${Account}:worker-configuration/${WorkerConfigurationName}/${UUID}"},
	})
}
