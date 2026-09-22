# arn:aws:devicefarm:ap-northeast-1:111111111111:networkprofile:resource-id
output "devicefarm_networkprofile" {
  value = provider::arn::devicefarm_networkprofile("resource-id")
}
