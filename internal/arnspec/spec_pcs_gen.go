// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: pcs
// Source: https://servicereference.us-east-1.amazonaws.com/v1/pcs/pcs.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "pcs_cluster", Service: "pcs", Resource: "cluster", Template: "arn:${Partition}:pcs:${Region}:${Account}:cluster/${ClusterIdentifier}"},
		{Name: "pcs_computenodegroup", Service: "pcs", Resource: "computenodegroup", Template: "arn:${Partition}:pcs:${Region}:${Account}:cluster/${ClusterIdentifier}/computenodegroup/${ComputeNodeGroupIdentifier}"},
		{Name: "pcs_queue", Service: "pcs", Resource: "queue", Template: "arn:${Partition}:pcs:${Region}:${Account}:cluster/${ClusterIdentifier}/queue/${QueueIdentifier}"},
	})
}
