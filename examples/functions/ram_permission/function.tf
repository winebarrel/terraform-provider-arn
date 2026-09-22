# arn:aws:ram::111111111111:permission/resource-path
output "ram_permission" {
  value = provider::arn::ram_permission("resource-path")
}
