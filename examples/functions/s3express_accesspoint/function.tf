# arn:aws:s3express:ap-northeast-1:111111111111:accesspoint/access-point-name
output "s3express_accesspoint" {
  value = provider::arn::s3express_accesspoint("access-point-name")
}
