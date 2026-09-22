# arn:aws:dms:ap-northeast-1:111111111111:cert:*
output "dms_certificate" {
  value = provider::arn::dms_certificate()
}
