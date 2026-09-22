# arn:aws:fms:ap-northeast-1:111111111111:policy/id
output "fms_policy" {
  value = provider::arn::fms_policy("id")
}
