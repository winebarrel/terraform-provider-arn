# arn:aws:ivs:ap-northeast-1:111111111111:playback-key/resource-id
output "ivs_playback_key_pair" {
  value = provider::arn::ivs_playback_key_pair("resource-id")
}
