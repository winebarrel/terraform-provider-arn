# arn:aws:kms:ap-northeast-1:111111111111:key/key-id
output "events_key" {
  value = provider::arn::events_key("key-id")
}
