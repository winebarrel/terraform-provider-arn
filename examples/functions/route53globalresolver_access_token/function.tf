# arn:aws:route53globalresolver::111111111111:access-token/id
output "route53globalresolver_access_token" {
  value = provider::arn::route53globalresolver_access_token("id")
}
