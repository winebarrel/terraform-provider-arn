# arn:aws:s3express:ap-northeast-1:111111111111:bucket/bucket-name
output "s3express_bucket" {
  value = provider::arn::s3express_bucket("bucket-name")
}
