// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: cloudhsm
// Source: https://servicereference.us-east-1.amazonaws.com/v1/cloudhsm/cloudhsm.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "cloudhsm_backup", Service: "cloudhsm", Resource: "backup", Template: "arn:${Partition}:cloudhsm:${Region}:${Account}:backup/${CloudHsmBackupInstanceName}"},
		{Name: "cloudhsm_cluster", Service: "cloudhsm", Resource: "cluster", Template: "arn:${Partition}:cloudhsm:${Region}:${Account}:cluster/${CloudHsmClusterInstanceName}"},
	})
}
