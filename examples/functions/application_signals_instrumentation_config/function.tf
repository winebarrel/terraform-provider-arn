# arn:aws:application-signals:ap-northeast-1:111111111111:instrumentationConfig/service/environment/signal-type/location-hash
output "application_signals_instrumentation_config" {
  value = provider::arn::application_signals_instrumentation_config("service", "environment", "signal-type", "location-hash")
}
