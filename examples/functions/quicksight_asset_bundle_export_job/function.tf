# arn:aws:quicksight:ap-northeast-1:111111111111:asset-bundle-export-job/resource-id
output "quicksight_asset_bundle_export_job" {
  value = provider::arn::quicksight_asset_bundle_export_job("resource-id")
}
