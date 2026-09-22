# arn:aws:route53globalresolver::111111111111:access-source/id
output "route53globalresolver_access_source" {
  value = provider::arn::route53globalresolver_access_source("id")
}
