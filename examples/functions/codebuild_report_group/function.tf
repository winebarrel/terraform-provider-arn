# arn:aws:codebuild:ap-northeast-1:111111111111:report-group/report-group-name
output "codebuild_report_group" {
  value = provider::arn::codebuild_report_group("report-group-name")
}
