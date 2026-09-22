# arn:aws:license-manager:ap-northeast-1:111111111111:license-configuration:license-configuration-id
output "license_manager_license_configuration" {
  value = provider::arn::license_manager_license_configuration("license-configuration-id")
}
