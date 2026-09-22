# arn:aws:imagebuilder:ap-northeast-1:111111111111:infrastructure-configuration/resource-id
output "imagebuilder_infrastructure_configuration" {
  value = provider::arn::imagebuilder_infrastructure_configuration("resource-id")
}
