# arn:aws:iotsitewise:ap-northeast-1:111111111111:asset/asset-id
output "iotsitewise_asset" {
  value = provider::arn::iotsitewise_asset("asset-id")
}
