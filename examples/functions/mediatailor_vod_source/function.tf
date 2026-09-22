# arn:aws:mediatailor:ap-northeast-1:111111111111:vodSource/source-location-name/vod-source-name
output "mediatailor_vod_source" {
  value = provider::arn::mediatailor_vod_source("source-location-name", "vod-source-name")
}
