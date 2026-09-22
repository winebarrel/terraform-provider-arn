# arn:aws:devicefarm:ap-northeast-1:111111111111:vpceconfiguration:resource-id
output "devicefarm_vpceconfiguration" {
  value = provider::arn::devicefarm_vpceconfiguration("resource-id")
}
