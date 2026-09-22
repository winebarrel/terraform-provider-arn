# arn:aws:license-manager::111111111111:grant:grant-id
output "license_manager_grant" {
  value = provider::arn::license_manager_grant("grant-id")
}
