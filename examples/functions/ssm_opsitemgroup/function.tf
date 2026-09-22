# arn:aws:ssm:ap-northeast-1:111111111111:opsitemgroup/default
output "ssm_opsitemgroup" {
  value = provider::arn::ssm_opsitemgroup()
}
