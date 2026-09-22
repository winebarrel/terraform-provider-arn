# arn:aws:quicksight:ap-northeast-1:111111111111:action-connector/resource-id
output "quicksight_actionconnector" {
  value = provider::arn::quicksight_actionconnector("resource-id")
}
