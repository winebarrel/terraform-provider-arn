# arn:aws:wellarchitected:ap-northeast-1:111111111111:profile/resource-id
output "wellarchitected_profile" {
  value = provider::arn::wellarchitected_profile("resource-id")
}
