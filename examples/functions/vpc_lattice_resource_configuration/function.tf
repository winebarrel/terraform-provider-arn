# arn:aws:vpc-lattice:ap-northeast-1:111111111111:resourceconfiguration/resource-configuration-id
output "vpc_lattice_resource_configuration" {
  value = provider::arn::vpc_lattice_resource_configuration("resource-configuration-id")
}
