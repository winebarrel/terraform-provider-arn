# arn:aws:ivs:ap-northeast-1:111111111111:playback-restriction-policy/resource-id
output "ivs_playback_restriction_policy" {
  value = provider::arn::ivs_playback_restriction_policy("resource-id")
}
