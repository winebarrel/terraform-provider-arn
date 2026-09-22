# arn:aws:dms:ap-northeast-1:111111111111:data-provider:*
output "dms_data_provider" {
  value = provider::arn::dms_data_provider()
}
