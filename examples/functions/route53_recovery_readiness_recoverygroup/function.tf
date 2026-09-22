# arn:aws:route53-recovery-readiness::111111111111:recovery-group/resource-id
output "route53_recovery_readiness_recoverygroup" {
  value = provider::arn::route53_recovery_readiness_recoverygroup("resource-id")
}
