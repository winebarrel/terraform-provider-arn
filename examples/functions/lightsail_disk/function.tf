# arn:aws:lightsail:ap-northeast-1:111111111111:Disk/id
output "lightsail_disk" {
  value = provider::arn::lightsail_disk("id")
}
