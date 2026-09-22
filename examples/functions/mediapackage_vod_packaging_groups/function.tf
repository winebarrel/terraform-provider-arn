# arn:aws:mediapackage-vod:ap-northeast-1:111111111111:packaging-groups/packaging-group-identifier
output "mediapackage_vod_packaging_groups" {
  value = provider::arn::mediapackage_vod_packaging_groups("packaging-group-identifier")
}
