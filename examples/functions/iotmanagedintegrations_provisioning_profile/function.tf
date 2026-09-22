# arn:aws:iotmanagedintegrations:ap-northeast-1:111111111111:provisioning-profile/identifier
output "iotmanagedintegrations_provisioning_profile" {
  value = provider::arn::iotmanagedintegrations_provisioning_profile("identifier")
}
