// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: controlcatalog
// Source: https://servicereference.us-east-1.amazonaws.com/v1/controlcatalog/controlcatalog.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "controlcatalog_common_control", Service: "controlcatalog", Resource: "common-control", Template: "arn:${Partition}:controlcatalog:::common-control/${CommonControlId}"},
		{Name: "controlcatalog_control", Service: "controlcatalog", Resource: "control", Template: "arn:${Partition}:controlcatalog:::control/${ControlId}"},
		{Name: "controlcatalog_domain", Service: "controlcatalog", Resource: "domain", Template: "arn:${Partition}:controlcatalog:::domain/${DomainId}"},
		{Name: "controlcatalog_objective", Service: "controlcatalog", Resource: "objective", Template: "arn:${Partition}:controlcatalog:::objective/${ObjectiveId}"},
	})
}
