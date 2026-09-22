# arn:aws:mobiletargeting:ap-northeast-1:111111111111:reports
output "mobiletargeting_reports" {
  value = provider::arn::mobiletargeting_reports()
}
