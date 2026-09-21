// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: iotdeviceadvisor
// Source: https://servicereference.us-east-1.amazonaws.com/v1/iotdeviceadvisor/iotdeviceadvisor.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "iotdeviceadvisor_suitedefinition", Service: "iotdeviceadvisor", Resource: "Suitedefinition", Template: "arn:${Partition}:iotdeviceadvisor:${Region}:${Account}:suitedefinition/${SuiteDefinitionId}"},
		{Name: "iotdeviceadvisor_suiterun", Service: "iotdeviceadvisor", Resource: "Suiterun", Template: "arn:${Partition}:iotdeviceadvisor:${Region}:${Account}:suiterun/${SuiteDefinitionId}/${SuiteRunId}"},
	})
}
