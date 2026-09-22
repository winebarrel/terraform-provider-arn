# arn:aws:quicksight:ap-northeast-1:111111111111:analysis/resource-id
output "quicksight_analysis" {
  value = provider::arn::quicksight_analysis("resource-id")
}
