# arn:aws:personalize:ap-northeast-1:111111111111:campaign/resource-id
output "personalize_campaign" {
  value = provider::arn::personalize_campaign("resource-id")
}
