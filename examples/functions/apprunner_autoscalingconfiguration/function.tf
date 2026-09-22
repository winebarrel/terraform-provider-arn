# arn:aws:apprunner:ap-northeast-1:111111111111:autoscalingconfiguration/autoscaling-configuration-name/autoscaling-configuration-version/autoscaling-configuration-id
output "apprunner_autoscalingconfiguration" {
  value = provider::arn::apprunner_autoscalingconfiguration("autoscaling-configuration-name", "autoscaling-configuration-version", "autoscaling-configuration-id")
}
