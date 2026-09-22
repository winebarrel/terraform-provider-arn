# arn:aws:lightsail:ap-northeast-1:111111111111:Instance/id
output "lightsail_instance" {
  value = provider::arn::lightsail_instance("id")
}
