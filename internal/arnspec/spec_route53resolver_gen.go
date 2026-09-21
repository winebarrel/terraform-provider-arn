// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: route53resolver
// Source: https://servicereference.us-east-1.amazonaws.com/v1/route53resolver/route53resolver.json
// Functions: 11
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "route53resolver_autodefined_rule", Service: "route53resolver", Resource: "autodefined-rule", Template: "arn:${Partition}:route53resolver:${Region}:${Account}:autodefined-rule/${ResourceId}"},
		{Name: "route53resolver_firewall_config", Service: "route53resolver", Resource: "firewall-config", Template: "arn:${Partition}:route53resolver:${Region}:${Account}:firewall-config/${ResourceId}"},
		{Name: "route53resolver_firewall_domain_list", Service: "route53resolver", Resource: "firewall-domain-list", Template: "arn:${Partition}:route53resolver:${Region}:${Account}:firewall-domain-list/${ResourceId}"},
		{Name: "route53resolver_firewall_rule_group", Service: "route53resolver", Resource: "firewall-rule-group", Template: "arn:${Partition}:route53resolver:${Region}:${Account}:firewall-rule-group/${ResourceId}"},
		{Name: "route53resolver_firewall_rule_group_association", Service: "route53resolver", Resource: "firewall-rule-group-association", Template: "arn:${Partition}:route53resolver:${Region}:${Account}:firewall-rule-group-association/${ResourceId}"},
		{Name: "route53resolver_outpost_resolver", Service: "route53resolver", Resource: "outpost-resolver", Template: "arn:${Partition}:route53resolver:${Region}:${Account}:outpost-resolver/${ResourceId}"},
		{Name: "route53resolver_resolver_config", Service: "route53resolver", Resource: "resolver-config", Template: "arn:${Partition}:route53resolver:${Region}:${Account}:resolver-config/${ResourceId}"},
		{Name: "route53resolver_resolver_dnssec_config", Service: "route53resolver", Resource: "resolver-dnssec-config", Template: "arn:${Partition}:route53resolver:${Region}:${Account}:resolver-dnssec-config/${ResourceId}"},
		{Name: "route53resolver_resolver_endpoint", Service: "route53resolver", Resource: "resolver-endpoint", Template: "arn:${Partition}:route53resolver:${Region}:${Account}:resolver-endpoint/${ResourceId}"},
		{Name: "route53resolver_resolver_query_log_config", Service: "route53resolver", Resource: "resolver-query-log-config", Template: "arn:${Partition}:route53resolver:${Region}:${Account}:resolver-query-log-config/${ResourceId}"},
		{Name: "route53resolver_resolver_rule", Service: "route53resolver", Resource: "resolver-rule", Template: "arn:${Partition}:route53resolver:${Region}:${Account}:resolver-rule/${ResourceId}"},
	})
}
