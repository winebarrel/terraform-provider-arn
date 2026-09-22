# arn:aws:lightsail:ap-northeast-1:111111111111:Domain/id
output "lightsail_domain" {
  value = provider::arn::lightsail_domain("id")
}
