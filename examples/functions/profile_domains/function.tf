# arn:aws:profile:ap-northeast-1:111111111111:domains/domain-name
output "profile_domains" {
  value = provider::arn::profile_domains("domain-name")
}
