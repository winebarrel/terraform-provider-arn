# arn:aws:route53-recovery-readiness::111111111111:resource-set/resource-id
output "route53_recovery_readiness_resourceset" {
  value = provider::arn::route53_recovery_readiness_resourceset("resource-id")
}
