# arn:aws:vpc-lattice:ap-northeast-1:111111111111:servicenetworkresourceassociation/service-network-resource-association-id
output "vpc_lattice_service_network_resource_association" {
  value = provider::arn::vpc_lattice_service_network_resource_association("service-network-resource-association-id")
}
