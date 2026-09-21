// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: codeguru-security
// Source: https://servicereference.us-east-1.amazonaws.com/v1/codeguru-security/codeguru-security.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "codeguru_security_scan_name", Service: "codeguru-security", Resource: "ScanName", Template: "arn:${Partition}:codeguru-security:${Region}:${Account}:scans/${ScanName}"},
	})
}
