# arn:aws:s3:ap-northeast-1:111111111111:access-grants/default/location/token
output "s3_accessgrantslocation" {
  value = provider::arn::s3_accessgrantslocation("token")
}
