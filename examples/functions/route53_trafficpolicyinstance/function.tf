# arn:aws:route53:::trafficpolicyinstance/id
output "route53_trafficpolicyinstance" {
  value = provider::arn::route53_trafficpolicyinstance("id")
}
