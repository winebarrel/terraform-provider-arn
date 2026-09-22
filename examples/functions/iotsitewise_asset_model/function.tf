# arn:aws:iotsitewise:ap-northeast-1:111111111111:asset-model/asset-model-id
output "iotsitewise_asset_model" {
  value = provider::arn::iotsitewise_asset_model("asset-model-id")
}
