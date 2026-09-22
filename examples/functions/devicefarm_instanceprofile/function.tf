# arn:aws:devicefarm:ap-northeast-1:111111111111:instanceprofile:resource-id
output "devicefarm_instanceprofile" {
  value = provider::arn::devicefarm_instanceprofile("resource-id")
}
