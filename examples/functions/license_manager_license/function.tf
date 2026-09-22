# arn:aws:license-manager::111111111111:license:license-id
output "license_manager_license" {
  value = provider::arn::license_manager_license("license-id")
}
