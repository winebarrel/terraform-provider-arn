// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: lambda
// Source: https://servicereference.us-east-1.amazonaws.com/v1/lambda/lambda.json
// Functions: 11
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "lambda_capacity_provider", Service: "lambda", Resource: "capacityProvider", Template: "arn:${Partition}:lambda:${Region}:${Account}:capacity-provider:${CapacityProviderName}"},
		{Name: "lambda_code_signing_config", Service: "lambda", Resource: "code signing config", Template: "arn:${Partition}:lambda:${Region}:${Account}:code-signing-config:${CodeSigningConfigId}"},
		{Name: "lambda_durable_execution", Service: "lambda", Resource: "durable execution", Template: "arn:${Partition}:lambda:${Region}:${Account}:function:${FunctionName}:${Version}/durable-execution/${ExecutionName}/${ExecutionId}"},
		{Name: "lambda_event_source_mapping", Service: "lambda", Resource: "eventSourceMapping", Template: "arn:${Partition}:lambda:${Region}:${Account}:event-source-mapping:${UUID}"},
		{Name: "lambda_function", Service: "lambda", Resource: "function", Template: "arn:${Partition}:lambda:${Region}:${Account}:function:${FunctionName}"},
		{Name: "lambda_function_alias", Service: "lambda", Resource: "function alias", Template: "arn:${Partition}:lambda:${Region}:${Account}:function:${FunctionName}:${Alias}"},
		{Name: "lambda_function_version", Service: "lambda", Resource: "function version", Template: "arn:${Partition}:lambda:${Region}:${Account}:function:${FunctionName}:${Version}"},
		{Name: "lambda_layer", Service: "lambda", Resource: "layer", Template: "arn:${Partition}:lambda:${Region}:${Account}:layer:${LayerName}"},
		{Name: "lambda_layer_version", Service: "lambda", Resource: "layerVersion", Template: "arn:${Partition}:lambda:${Region}:${Account}:layer:${LayerName}:${LayerVersion}"},
		{Name: "lambda_microvm_image", Service: "lambda", Resource: "microvmImage", Template: "arn:${Partition}:lambda:${Region}:${Account}:microvm-image:${MicrovmImageName}"},
		{Name: "lambda_network_connector", Service: "lambda", Resource: "networkConnector", Template: "arn:${Partition}:lambda:${Region}:${Account}:network-connector:${NetworkConnectorId}"},
	})
}
