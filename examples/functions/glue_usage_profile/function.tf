# arn:aws:glue:ap-northeast-1:111111111111:usageProfile/usage-profile-id
output "glue_usage_profile" {
  value = provider::arn::glue_usage_profile("usage-profile-id")
}
