# arn:aws:kms:ap-northeast-1:111111111111:alias/alias
output "kms_alias" {
  value = provider::arn::kms_alias("alias")
}
