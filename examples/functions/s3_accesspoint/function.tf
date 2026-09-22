# arn:aws:s3:ap-northeast-1:111111111111:accesspoint/access-point-name
output "s3_accesspoint" {
  value = provider::arn::s3_accesspoint("access-point-name")
}
