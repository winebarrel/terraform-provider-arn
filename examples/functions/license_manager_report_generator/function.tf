# arn:aws:license-manager:ap-northeast-1:111111111111:report-generator:report-generator-id
output "license_manager_report_generator" {
  value = provider::arn::license_manager_report_generator("report-generator-id")
}
