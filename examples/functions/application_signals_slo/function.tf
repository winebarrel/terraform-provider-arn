# arn:aws:application-signals:ap-northeast-1:111111111111:slo/slo-name
output "application_signals_slo" {
  value = provider::arn::application_signals_slo("slo-name")
}
