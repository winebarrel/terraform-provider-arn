# arn:aws:vpc-lattice:ap-northeast-1:111111111111:servicenetwork/service-network-id
output "vpc_lattice_service_network" {
  value = provider::arn::vpc_lattice_service_network("service-network-id")
}
