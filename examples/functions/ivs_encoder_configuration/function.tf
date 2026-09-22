# arn:aws:ivs:ap-northeast-1:111111111111:encoder-configuration/resource-id
output "ivs_encoder_configuration" {
  value = provider::arn::ivs_encoder_configuration("resource-id")
}
