# arn:aws:ssm:ap-northeast-1:111111111111:managed-instance-inventory/instance-id
output "ssm_managed_instance_inventory" {
  value = provider::arn::ssm_managed_instance_inventory("instance-id")
}
