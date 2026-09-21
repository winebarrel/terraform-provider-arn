// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: route53-recovery-cluster
// Source: https://servicereference.us-east-1.amazonaws.com/v1/route53-recovery-cluster/route53-recovery-cluster.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "route53_recovery_cluster_routingcontrol", Service: "route53-recovery-cluster", Resource: "routingcontrol", Template: "arn:${Partition}:route53-recovery-control::${Account}:controlpanel/${ControlPanelId}/routingcontrol/${RoutingControlId}"},
	})
}
