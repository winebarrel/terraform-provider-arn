// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: ebs
// Source: https://servicereference.us-east-1.amazonaws.com/v1/ebs/ebs.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "ebs_snapshot", Service: "ebs", Resource: "snapshot", Template: "arn:${Partition}:ec2:${Region}::snapshot/${SnapshotId}"},
	})
}
