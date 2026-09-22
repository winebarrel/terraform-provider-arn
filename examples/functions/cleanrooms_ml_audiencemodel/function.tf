# arn:aws:cleanrooms-ml:ap-northeast-1:111111111111:audience-model/resource-id
output "cleanrooms_ml_audiencemodel" {
  value = provider::arn::cleanrooms_ml_audiencemodel("resource-id")
}
