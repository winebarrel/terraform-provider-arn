# arn:aws:s3:ap-northeast-1:111111111111:access-grants/default
output "s3_accessgrantsinstance" {
  value = provider::arn::s3_accessgrantsinstance()
}
