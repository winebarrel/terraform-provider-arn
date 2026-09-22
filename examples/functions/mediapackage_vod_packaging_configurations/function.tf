# arn:aws:mediapackage-vod:ap-northeast-1:111111111111:packaging-configurations/packaging-configuration-identifier
output "mediapackage_vod_packaging_configurations" {
  value = provider::arn::mediapackage_vod_packaging_configurations("packaging-configuration-identifier")
}
