# arn:aws:vpc-lattice:ap-northeast-1:111111111111:resourcegateway/resource-gateway-id
output "vpc_lattice_resource_gateway" {
  value = provider::arn::vpc_lattice_resource_gateway("resource-gateway-id")
}
