# arn:aws:route53:::change/id
output "route53_change" {
  value = provider::arn::route53_change("id")
}
