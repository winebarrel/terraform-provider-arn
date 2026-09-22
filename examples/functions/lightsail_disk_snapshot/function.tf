# arn:aws:lightsail:ap-northeast-1:111111111111:DiskSnapshot/id
output "lightsail_disk_snapshot" {
  value = provider::arn::lightsail_disk_snapshot("id")
}
