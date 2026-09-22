# arn:aws:ssm:ap-northeast-1:111111111111:opsmetadata/resource-id
output "ssm_opsmetadata" {
  value = provider::arn::ssm_opsmetadata("resource-id")
}
