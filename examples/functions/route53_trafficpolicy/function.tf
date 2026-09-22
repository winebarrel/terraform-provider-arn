# arn:aws:route53:::trafficpolicy/id
output "route53_trafficpolicy" {
  value = provider::arn::route53_trafficpolicy("id")
}
