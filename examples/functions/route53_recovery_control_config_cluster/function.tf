# arn:aws:route53-recovery-control::111111111111:cluster/resource-id
output "route53_recovery_control_config_cluster" {
  value = provider::arn::route53_recovery_control_config_cluster("resource-id")
}
