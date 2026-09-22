# arn:aws:macie2:ap-northeast-1:111111111111:findings-filter/resource-id
output "macie2_findings_filter" {
  value = provider::arn::macie2_findings_filter("resource-id")
}
