// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: arc-zonal-shift
// Source: https://servicereference.us-east-1.amazonaws.com/v1/arc-zonal-shift/arc-zonal-shift.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "arc_zonal_shift_alb", Service: "arc-zonal-shift", Resource: "ALB", Template: "arn:${Partition}:elasticloadbalancing:${Region}:${Account}:loadbalancer/app/${LoadBalancerName}/${LoadBalancerId}"},
		{Name: "arc_zonal_shift_nlb", Service: "arc-zonal-shift", Resource: "NLB", Template: "arn:${Partition}:elasticloadbalancing:${Region}:${Account}:loadbalancer/net/${LoadBalancerName}/${LoadBalancerId}"},
	})
}
