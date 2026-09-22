# arn:aws:lightsail:ap-northeast-1:111111111111:InstanceSnapshot/id
output "lightsail_instance_snapshot" {
  value = provider::arn::lightsail_instance_snapshot("id")
}
