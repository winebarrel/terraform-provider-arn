# arn:aws:ivs:ap-northeast-1:111111111111:public-key/resource-id
output "ivs_public_key" {
  value = provider::arn::ivs_public_key("resource-id")
}
