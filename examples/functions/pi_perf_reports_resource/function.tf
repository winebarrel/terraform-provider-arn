# arn:aws:pi:ap-northeast-1:111111111111:perf-reports/service-type/identifier/report-id
output "pi_perf_reports_resource" {
  value = provider::arn::pi_perf_reports_resource("service-type", "identifier", "report-id")
}
