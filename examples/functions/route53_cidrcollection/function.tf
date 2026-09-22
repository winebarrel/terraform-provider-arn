# arn:aws:route53:::cidrcollection/id
output "route53_cidrcollection" {
  value = provider::arn::route53_cidrcollection("id")
}
