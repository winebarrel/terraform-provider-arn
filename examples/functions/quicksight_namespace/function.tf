# arn:aws:quicksight:ap-northeast-1:111111111111:namespace/resource-id
output "quicksight_namespace" {
  value = provider::arn::quicksight_namespace("resource-id")
}
