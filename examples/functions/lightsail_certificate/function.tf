# arn:aws:lightsail:ap-northeast-1:111111111111:Certificate/id
output "lightsail_certificate" {
  value = provider::arn::lightsail_certificate("id")
}
