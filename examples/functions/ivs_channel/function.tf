# arn:aws:ivs:ap-northeast-1:111111111111:channel/resource-id
output "ivs_channel" {
  value = provider::arn::ivs_channel("resource-id")
}
