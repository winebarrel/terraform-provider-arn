// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: route53globalresolver
// Source: https://servicereference.us-east-1.amazonaws.com/v1/route53globalresolver/route53globalresolver.json
// Functions: 5
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "route53globalresolver_access_source", Service: "route53globalresolver", Resource: "access-source", Template: "arn:${Partition}:route53globalresolver::${Account}:access-source/${Id}"},
		{Name: "route53globalresolver_access_token", Service: "route53globalresolver", Resource: "access-token", Template: "arn:${Partition}:route53globalresolver::${Account}:access-token/${Id}"},
		{Name: "route53globalresolver_dns_view", Service: "route53globalresolver", Resource: "dns-view", Template: "arn:${Partition}:route53globalresolver::${Account}:dns-view/${Id}"},
		{Name: "route53globalresolver_firewall_domain_list", Service: "route53globalresolver", Resource: "firewall-domain-list", Template: "arn:${Partition}:route53globalresolver::${Account}:firewall-domain-list/${Id}"},
		{Name: "route53globalresolver_global_resolver", Service: "route53globalresolver", Resource: "global-resolver", Template: "arn:${Partition}:route53globalresolver::${Account}:global-resolver/${Id}"},
	})
}
