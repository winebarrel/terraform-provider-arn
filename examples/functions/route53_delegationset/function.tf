# arn:aws:route53:::delegationset/id
output "route53_delegationset" {
  value = provider::arn::route53_delegationset("id")
}
