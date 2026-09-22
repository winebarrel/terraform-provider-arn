# arn:aws:wellarchitected:ap-northeast-1:111111111111:lens/resource-id
output "wellarchitected_lens" {
  value = provider::arn::wellarchitected_lens("resource-id")
}
