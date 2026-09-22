# arn:aws:kms:ap-northeast-1:111111111111:key/key-id
output "kinesis_kms_key" {
  value = provider::arn::kinesis_kms_key("key-id")
}
