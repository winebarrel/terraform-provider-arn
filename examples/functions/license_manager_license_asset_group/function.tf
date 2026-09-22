# arn:aws:license-manager:ap-northeast-1:111111111111:license-asset-group:license-asset-group-id
output "license_manager_license_asset_group" {
  value = provider::arn::license_manager_license_asset_group("license-asset-group-id")
}
