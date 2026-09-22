# arn:aws:quicksight:ap-northeast-1:111111111111:asset-bundle-import-job/resource-id
output "quicksight_asset_bundle_import_job" {
  value = provider::arn::quicksight_asset_bundle_import_job("resource-id")
}
