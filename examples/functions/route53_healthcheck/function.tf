# arn:aws:route53:::healthcheck/id
output "route53_healthcheck" {
  value = provider::arn::route53_healthcheck("id")
}
