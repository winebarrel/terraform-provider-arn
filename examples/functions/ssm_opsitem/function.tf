# arn:aws:ssm:ap-northeast-1:111111111111:opsitem/resource-id
output "ssm_opsitem" {
  value = provider::arn::ssm_opsitem("resource-id")
}
