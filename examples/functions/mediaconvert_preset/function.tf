# arn:aws:mediaconvert:ap-northeast-1:111111111111:presets/preset-name
output "mediaconvert_preset" {
  value = provider::arn::mediaconvert_preset("preset-name")
}
