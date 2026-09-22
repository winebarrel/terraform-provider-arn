# arn:aws:dms:ap-northeast-1:111111111111:endpoint:*
output "dms_endpoint" {
  value = provider::arn::dms_endpoint()
}
