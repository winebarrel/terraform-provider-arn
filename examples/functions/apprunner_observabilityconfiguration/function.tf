# arn:aws:apprunner:ap-northeast-1:111111111111:observabilityconfiguration/observability-configuration-name/observability-configuration-version/observability-configuration-id
output "apprunner_observabilityconfiguration" {
  value = provider::arn::apprunner_observabilityconfiguration("observability-configuration-name", "observability-configuration-version", "observability-configuration-id")
}
