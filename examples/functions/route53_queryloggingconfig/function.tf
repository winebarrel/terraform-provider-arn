# arn:aws:route53:::queryloggingconfig/id
output "route53_queryloggingconfig" {
  value = provider::arn::route53_queryloggingconfig("id")
}
