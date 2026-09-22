# arn:aws:route53globalresolver::111111111111:global-resolver/id
output "route53globalresolver_global_resolver" {
  value = provider::arn::route53globalresolver_global_resolver("id")
}
