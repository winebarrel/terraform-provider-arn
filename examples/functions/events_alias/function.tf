# arn:aws:kms:ap-northeast-1:111111111111:alias/alias
output "events_alias" {
  value = provider::arn::events_alias("alias")
}
