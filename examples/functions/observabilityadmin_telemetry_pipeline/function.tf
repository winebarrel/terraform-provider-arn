# arn:aws:observabilityadmin:ap-northeast-1:111111111111:telemetry-pipeline/telemetry-pipeline-identifier
output "observabilityadmin_telemetry_pipeline" {
  value = provider::arn::observabilityadmin_telemetry_pipeline("telemetry-pipeline-identifier")
}
