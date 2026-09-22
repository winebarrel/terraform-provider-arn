# arn:aws:route53:::hostedzone/id
output "route53_hostedzone" {
  value = provider::arn::route53_hostedzone("id")
}
