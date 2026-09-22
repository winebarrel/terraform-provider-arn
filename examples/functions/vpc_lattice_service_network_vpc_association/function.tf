# arn:aws:vpc-lattice:ap-northeast-1:111111111111:servicenetworkvpcassociation/service-network-vpc-association-id
output "vpc_lattice_service_network_vpc_association" {
  value = provider::arn::vpc_lattice_service_network_vpc_association("service-network-vpc-association-id")
}
