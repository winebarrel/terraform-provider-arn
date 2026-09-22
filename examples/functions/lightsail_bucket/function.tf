# arn:aws:lightsail:ap-northeast-1:111111111111:Bucket/id
output "lightsail_bucket" {
  value = provider::arn::lightsail_bucket("id")
}
