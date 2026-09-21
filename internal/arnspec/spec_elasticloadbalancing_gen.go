// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: elasticloadbalancing
// Source: https://servicereference.us-east-1.amazonaws.com/v1/elasticloadbalancing/elasticloadbalancing.json
// Functions: 11
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "elasticloadbalancing_listener_app", Service: "elasticloadbalancing", Resource: "listener/app", Template: "arn:${Partition}:elasticloadbalancing:${Region}:${Account}:listener/app/${LoadBalancerName}/${LoadBalancerId}/${ListenerId}"},
		{Name: "elasticloadbalancing_listener_gwy", Service: "elasticloadbalancing", Resource: "listener/gwy", Template: "arn:${Partition}:elasticloadbalancing:${Region}:${Account}:listener/gwy/${LoadBalancerName}/${LoadBalancerId}/${ListenerId}"},
		{Name: "elasticloadbalancing_listener_net", Service: "elasticloadbalancing", Resource: "listener/net", Template: "arn:${Partition}:elasticloadbalancing:${Region}:${Account}:listener/net/${LoadBalancerName}/${LoadBalancerId}/${ListenerId}"},
		{Name: "elasticloadbalancing_listener_rule_app", Service: "elasticloadbalancing", Resource: "listener-rule/app", Template: "arn:${Partition}:elasticloadbalancing:${Region}:${Account}:listener-rule/app/${LoadBalancerName}/${LoadBalancerId}/${ListenerId}/${ListenerRuleId}"},
		{Name: "elasticloadbalancing_listener_rule_net", Service: "elasticloadbalancing", Resource: "listener-rule/net", Template: "arn:${Partition}:elasticloadbalancing:${Region}:${Account}:listener-rule/net/${LoadBalancerName}/${LoadBalancerId}/${ListenerId}/${ListenerRuleId}"},
		{Name: "elasticloadbalancing_loadbalancer", Service: "elasticloadbalancing", Resource: "loadbalancer", Template: "arn:${Partition}:elasticloadbalancing:${Region}:${Account}:loadbalancer/${LoadBalancerName}"},
		{Name: "elasticloadbalancing_loadbalancer_app", Service: "elasticloadbalancing", Resource: "loadbalancer/app/", Template: "arn:${Partition}:elasticloadbalancing:${Region}:${Account}:loadbalancer/app/${LoadBalancerName}/${LoadBalancerId}"},
		{Name: "elasticloadbalancing_loadbalancer_gwy", Service: "elasticloadbalancing", Resource: "loadbalancer/gwy/", Template: "arn:${Partition}:elasticloadbalancing:${Region}:${Account}:loadbalancer/gwy/${LoadBalancerName}/${LoadBalancerId}"},
		{Name: "elasticloadbalancing_loadbalancer_net", Service: "elasticloadbalancing", Resource: "loadbalancer/net/", Template: "arn:${Partition}:elasticloadbalancing:${Region}:${Account}:loadbalancer/net/${LoadBalancerName}/${LoadBalancerId}"},
		{Name: "elasticloadbalancing_targetgroup", Service: "elasticloadbalancing", Resource: "targetgroup", Template: "arn:${Partition}:elasticloadbalancing:${Region}:${Account}:targetgroup/${TargetGroupName}/${TargetGroupId}"},
		{Name: "elasticloadbalancing_truststore", Service: "elasticloadbalancing", Resource: "truststore", Template: "arn:${Partition}:elasticloadbalancing:${Region}:${Account}:truststore/${TrustStoreName}/${TrustStoreId}"},
	})
}
