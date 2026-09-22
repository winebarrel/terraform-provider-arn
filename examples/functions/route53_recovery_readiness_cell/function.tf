# arn:aws:route53-recovery-readiness::111111111111:cell/resource-id
output "route53_recovery_readiness_cell" {
  value = provider::arn::route53_recovery_readiness_cell("resource-id")
}
