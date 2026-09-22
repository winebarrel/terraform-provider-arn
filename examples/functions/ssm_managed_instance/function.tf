# arn:aws:ssm:ap-northeast-1:111111111111:managed-instance/instance-id
output "ssm_managed_instance" {
  value = provider::arn::ssm_managed_instance("instance-id")
}
