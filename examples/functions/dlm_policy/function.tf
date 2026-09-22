# arn:aws:dlm:ap-northeast-1:111111111111:policy/resource-name
output "dlm_policy" {
  value = provider::arn::dlm_policy("resource-name")
}
