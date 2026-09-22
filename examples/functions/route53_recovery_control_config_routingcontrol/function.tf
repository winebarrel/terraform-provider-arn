# arn:aws:route53-recovery-control::111111111111:controlpanel/control-panel-id/routingcontrol/routing-control-id
output "route53_recovery_control_config_routingcontrol" {
  value = provider::arn::route53_recovery_control_config_routingcontrol("control-panel-id", "routing-control-id")
}
