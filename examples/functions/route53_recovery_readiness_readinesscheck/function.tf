# arn:aws:route53-recovery-readiness::111111111111:readiness-check/resource-id
output "route53_recovery_readiness_readinesscheck" {
  value = provider::arn::route53_recovery_readiness_readinesscheck("resource-id")
}
