# arn:aws:iotmanagedintegrations:ap-northeast-1:111111111111:managed-thing/identifier
output "iotmanagedintegrations_managed_thing" {
  value = provider::arn::iotmanagedintegrations_managed_thing("identifier")
}
