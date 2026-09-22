# arn:aws:s3:ap-northeast-1:111111111111:access-grants/default/grant/token
output "s3_accessgrant" {
  value = provider::arn::s3_accessgrant("token")
}
