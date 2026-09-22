# arn:aws:route53-recovery-control::111111111111:controlpanel/control-panel-id
output "route53_recovery_control_config_controlpanel" {
  value = provider::arn::route53_recovery_control_config_controlpanel("control-panel-id")
}
