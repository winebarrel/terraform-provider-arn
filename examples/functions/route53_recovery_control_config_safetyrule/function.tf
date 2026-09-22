# arn:aws:route53-recovery-control::111111111111:controlpanel/control-panel-id/safetyrule/safety-rule-id
output "route53_recovery_control_config_safetyrule" {
  value = provider::arn::route53_recovery_control_config_safetyrule("control-panel-id", "safety-rule-id")
}
