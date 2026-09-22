# arn:aws:artifact:ap-northeast-1::report/report-id:version
output "artifact_report" {
  value = provider::arn::artifact_report("report-id", "version")
}
