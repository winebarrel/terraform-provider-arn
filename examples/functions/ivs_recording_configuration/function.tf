# arn:aws:ivs:ap-northeast-1:111111111111:recording-configuration/resource-id
output "ivs_recording_configuration" {
  value = provider::arn::ivs_recording_configuration("resource-id")
}
