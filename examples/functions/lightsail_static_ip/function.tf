# arn:aws:lightsail:ap-northeast-1:111111111111:StaticIp/id
output "lightsail_static_ip" {
  value = provider::arn::lightsail_static_ip("id")
}
