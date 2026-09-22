# arn:aws:kms:ap-northeast-1:111111111111:key/key-id
output "kms_key" {
  value = provider::arn::kms_key("key-id")
}
