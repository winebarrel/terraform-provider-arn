// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: appflow
// Source: https://servicereference.us-east-1.amazonaws.com/v1/appflow/appflow.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "appflow_connector", Service: "appflow", Resource: "connector", Template: "arn:${Partition}:appflow:${Region}:${Account}:connector/${ConnectorLabel}"},
		{Name: "appflow_connectorprofile", Service: "appflow", Resource: "connectorprofile", Template: "arn:${Partition}:appflow:${Region}:${Account}:connectorprofile/${ProfileName}"},
		{Name: "appflow_flow", Service: "appflow", Resource: "flow", Template: "arn:${Partition}:appflow:${Region}:${Account}:flow/${FlowName}"},
	})
}
