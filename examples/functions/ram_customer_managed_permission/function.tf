# arn:aws:ram:ap-northeast-1:111111111111:permission/resource-path
output "ram_customer_managed_permission" {
  value = provider::arn::ram_customer_managed_permission("resource-path")
}
