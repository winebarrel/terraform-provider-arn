# arn:aws:mediapackage-vod:ap-northeast-1:111111111111:assets/asset-identifier
output "mediapackage_vod_assets" {
  value = provider::arn::mediapackage_vod_assets("asset-identifier")
}
