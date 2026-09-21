// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: scn
// Source: https://servicereference.us-east-1.amazonaws.com/v1/scn/scn.json
// Functions: 5
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "scn_bill_of_materials_import_job", Service: "scn", Resource: "bill-of-materials-import-job", Template: "arn:${Partition}:scn:${Region}:${Account}:instance/${InstanceId}/bill-of-materials-import-job/${JobId}"},
		{Name: "scn_data_integration_flow", Service: "scn", Resource: "data-integration-flow", Template: "arn:${Partition}:scn:${Region}:${Account}:instance/${InstanceId}/data-integration-flows/${FlowName}"},
		{Name: "scn_dataset", Service: "scn", Resource: "dataset", Template: "arn:${Partition}:scn:${Region}:${Account}:instance/${InstanceId}/namespaces/${Namespace}/datasets/${DatasetName}"},
		{Name: "scn_instance", Service: "scn", Resource: "instance", Template: "arn:${Partition}:scn:${Region}:${Account}:instance/${InstanceId}"},
		{Name: "scn_namespace", Service: "scn", Resource: "namespace", Template: "arn:${Partition}:scn:${Region}:${Account}:instance/${InstanceId}/namespaces/${Namespace}"},
	})
}
