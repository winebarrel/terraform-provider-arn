# arn:aws:vpc-lattice:ap-northeast-1:111111111111:servicenetworkserviceassociation/service-network-service-association-id
output "vpc_lattice_service_network_service_association" {
  value = provider::arn::vpc_lattice_service_network_service_association("service-network-service-association-id")
}
