# arn:aws:route53globalresolver::111111111111:dns-view/id
output "route53globalresolver_dns_view" {
  value = provider::arn::route53globalresolver_dns_view("id")
}
