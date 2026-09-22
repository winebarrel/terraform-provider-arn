# arn:aws:quicksight:ap-northeast-1:111111111111:limits-profile/resource-id
output "quicksight_limits_profile" {
  value = provider::arn::quicksight_limits_profile("resource-id")
}
