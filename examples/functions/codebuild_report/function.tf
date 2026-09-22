# arn:aws:codebuild:ap-northeast-1:111111111111:report/report-group-name:report-id
output "codebuild_report" {
  value = provider::arn::codebuild_report("report-group-name", "report-id")
}
