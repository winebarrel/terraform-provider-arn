# arn:aws:route53globalresolver::111111111111:firewall-domain-list/id
output "route53globalresolver_firewall_domain_list" {
  value = provider::arn::route53globalresolver_firewall_domain_list("id")
}
