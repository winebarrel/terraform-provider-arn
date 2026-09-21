// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: network-firewall
// Source: https://servicereference.us-east-1.amazonaws.com/v1/network-firewall/network-firewall.json
// Functions: 10
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "network_firewall_container_association", Service: "network-firewall", Resource: "ContainerAssociation", Template: "arn:${Partition}:network-firewall:${Region}:${Account}:container-association/${Name}"},
		{Name: "network_firewall_firewall", Service: "network-firewall", Resource: "Firewall", Template: "arn:${Partition}:network-firewall:${Region}:${Account}:firewall/${Name}"},
		{Name: "network_firewall_firewall_policy", Service: "network-firewall", Resource: "FirewallPolicy", Template: "arn:${Partition}:network-firewall:${Region}:${Account}:firewall-policy/${Name}"},
		{Name: "network_firewall_proxy", Service: "network-firewall", Resource: "Proxy", Template: "arn:${Partition}:network-firewall:${Region}:${Account}:proxy/${Name}"},
		{Name: "network_firewall_proxy_configuration", Service: "network-firewall", Resource: "ProxyConfiguration", Template: "arn:${Partition}:network-firewall:${Region}:${Account}:proxy-configuration/${Name}"},
		{Name: "network_firewall_proxy_rule_group", Service: "network-firewall", Resource: "ProxyRuleGroup", Template: "arn:${Partition}:network-firewall:${Region}:${Account}:proxy-rule-group/${Name}"},
		{Name: "network_firewall_stateful_rule_group", Service: "network-firewall", Resource: "StatefulRuleGroup", Template: "arn:${Partition}:network-firewall:${Region}:${Account}:stateful-rulegroup/${Name}"},
		{Name: "network_firewall_stateless_rule_group", Service: "network-firewall", Resource: "StatelessRuleGroup", Template: "arn:${Partition}:network-firewall:${Region}:${Account}:stateless-rulegroup/${Name}"},
		{Name: "network_firewall_tls_inspection_configuration", Service: "network-firewall", Resource: "TLSInspectionConfiguration", Template: "arn:${Partition}:network-firewall:${Region}:${Account}:tls-configuration/${Name}"},
		{Name: "network_firewall_vpc_endpoint_association", Service: "network-firewall", Resource: "VpcEndpointAssociation", Template: "arn:${Partition}:network-firewall:${Region}:${Account}:vpc-endpoint-association/${Name}"},
	})
}
