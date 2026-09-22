# arn:aws:license-manager:ap-northeast-1:111111111111:license-asset-ruleset:license-asset-ruleset-id
output "license_manager_license_asset_ruleset" {
  value = provider::arn::license_manager_license_asset_ruleset("license-asset-ruleset-id")
}
