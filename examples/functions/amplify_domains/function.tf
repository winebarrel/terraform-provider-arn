# arn:aws:amplify:ap-northeast-1:111111111111:apps/app-id/domains/domain-name
output "amplify_domains" {
  value = provider::arn::amplify_domains("app-id", "domain-name")
}
