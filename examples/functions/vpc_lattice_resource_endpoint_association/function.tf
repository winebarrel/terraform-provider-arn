# arn:aws:vpc-lattice:ap-northeast-1:111111111111:resourceendpointassociation/resource-endpoint-association-id
output "vpc_lattice_resource_endpoint_association" {
  value = provider::arn::vpc_lattice_resource_endpoint_association("resource-endpoint-association-id")
}
