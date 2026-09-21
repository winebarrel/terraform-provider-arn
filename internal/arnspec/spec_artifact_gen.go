// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: artifact
// Source: https://servicereference.us-east-1.amazonaws.com/v1/artifact/artifact.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "artifact_agreement", Service: "artifact", Resource: "agreement", Template: "arn:${Partition}:artifact:::agreement/*"},
		{Name: "artifact_compliance_inquiry", Service: "artifact", Resource: "compliance-inquiry", Template: "arn:${Partition}:artifact:${Region}:${Account}:compliance-inquiry/*"},
		{Name: "artifact_customer_agreement", Service: "artifact", Resource: "customer-agreement", Template: "arn:${Partition}:artifact::${Account}:customer-agreement/*"},
		{Name: "artifact_report", Service: "artifact", Resource: "report", Template: "arn:${Partition}:artifact:${Region}::report/${ReportId}:${Version}"},
	})
}
