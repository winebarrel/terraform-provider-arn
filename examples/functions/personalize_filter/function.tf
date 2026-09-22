# arn:aws:personalize:ap-northeast-1:111111111111:filter/resource-id
output "personalize_filter" {
  value = provider::arn::personalize_filter("resource-id")
}
